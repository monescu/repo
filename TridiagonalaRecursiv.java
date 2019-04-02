package main;

import java.util.Scanner;

public class TridiagonalaRecursiv {

	static Scanner sc = new Scanner(System.in);
	static float[] x, d, a, b, r, c, s;
	static int n;

	public static void recrS( int i) {
		if (i == 1 && a[1] != 0) {
			r[1] = -c[1]/a[1];
			s[1] = d[1]/a[1];
		}
		if (i < n-1) {
			r[i+1] = -c[i+1] / (b[i] * r[i] + a[i+1]);
			s[i+1] = (d[i+1] - b[i] * s[i]) / (b[i] * r[i] + a[i+1]);
			recrS(i+1);
		}
	}

	public static void recxA(int i) {
		if (i == n) {
			x[i] = (d[i] - b[i-1] * s[i-1]) / (b[i-1] * r[i-1] + a[i]);
			recxA(i-1);
		} else if (i > 0) {
			x[i] = r[i] * x[i+1] + s[i];
			recxA(i-1);
		}
	}
	
	public static void main(String[] args) {
		
		for(int i=0; i<n; i++) {
			System.out.print("a["+i+"]= ");
			a[i] = sc.nextFloat();
		}
		for(int i=1; i<n; i++) {
			System.out.print("c["+i+"]= ");
			c[i] = sc.nextFloat();
		}
		for(int i=0; i<n-1; i++) {
			System.out.print("b["+i+"]= ");
			b[i] = sc.nextFloat();
		}
		for(int i=0; i<n; i++) {
			System.out.print("d["+i+"]= ");
			d[i] = sc.nextFloat();
		}
		r = new float[n];
		x = new float[n + 1];
		s = new float[n];
		
		recrS(1);
		recxA(n);
		
		System.out.println("\nSolutii:");
		for (int i = 1; i <= n; i++)
			System.out.println("x[" + i + "]: " + (float) Math.round(x[i] * 100000d) / 100000d);
	}
}