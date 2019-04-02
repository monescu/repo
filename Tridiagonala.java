package main;
import java.util.Scanner;
public class Tridiagonala {

	static Scanner sc = new Scanner(System.in);
	
	public static void main(String[] args) {
		int n;
		System.out.print("n= ");
		n = sc.nextInt();
		double a[] = new double [n]; 
		double b[] = new double [n];
		double c[] = new double [n];
		double d[] = new double [n];
		double x[] = new double [n];
		double r[] = new double [n];
		double s[] = new double [n];
		
		for(int i=0; i<n; i++) {
			System.out.print("a["+i+"]= ");
			a[i] = sc.nextDouble();
		}
		for(int i=1; i<n; i++) {
			System.out.print("c["+i+"]= ");
			c[i] = sc.nextDouble();
		}
		for(int i=0; i<n-1; i++) {
			System.out.print("b["+i+"]= ");
			b[i] = sc.nextDouble();
		}
		for(int i=0; i<n; i++) {
			System.out.print("d["+i+"]= ");
			d[i] = sc.nextDouble();
		}
		
		r[0] = -c[1] / a[0];
		s[0] = d[0] / a[0];
		
		for(int i=1; i<=n-2; i++) {
			r[i] = -c[i]/(b[i] * r[i-1] + a[i]);
			s[i] = (d[i] - b[i] * s[i-1])/(b[i] * r[i-1] + a[i]);
		}
		for(int i=1; i<n; i++) {
			x[i] = (d[i] - b[i] * s[i-1])/(b[i] * r[i-1] + a[i]);
		}
		
		x[0] = r[0] * x[1] + s[0];
		for(int i=0; i<n; i++)
			System.out.print(x[i]+" ");
	}

}
