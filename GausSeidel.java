package main;
import java.util.Scanner;

public class GausSeidel {

	static Scanner sc = new Scanner(System.in);
	
	private double[][] A;
	private double[] b;

	public GausSeidel(double[][] A, double[] b) {

		if (A == null || b == null)
			throw new NullPointerException();
		if (A.length != b.length)
			throw new IllegalArgumentException();
		this.A = A;
		this.b = b;

	}

	public boolean converge() {
		for (int i = 0; i < A.length; i++) {
			double diagonal = Math.abs(A[i][i]);
			double tmpSum = 0;
			for (int j = 0; j < A.length; j++)
				if (i != j)
					tmpSum += Math.abs(A[i][j]);
			if (tmpSum >= diagonal)
				return false;
		}
		return true;
	}

	public double[] rezolvare(int precision) {
		if (!converge())
			System.err.println("Solutia nu poate converge.");

		double[] x = init(new double[A.length]);
		for (int k = 0; k < precision; k++)
			for (int i = 0; i < A.length; i++) {
				double x0 = 0;
				for (int j = 0; j < A.length; j++)
					if (i != j)
						x0 += A[i][j] * x[j];
				x[i] = (b[i] - x0) / A[i][i];
			}
		return x;
	}

	private double[] init(double[] ds) {
		for (int i = 0; i < ds.length; i++)
			ds[i] = 0;
		return ds;
	}
	
	public static void main(String[] args) {
		int n;
		System.out.print("n= ");
		n = sc.nextInt();
		double [][]A = new double [n][n];
		double []b = new double [n];
		
		for(int i=0; i<n; i++)
			for(int j=0; j<n; j++)
				A[i][j] = sc.nextDouble();
	
		for(int i=0; i<n; i++) {
			System.out.print("b["+i+"]=");
			b[i] = sc.nextDouble();
		}
		GausSeidel g = new GausSeidel(A, b);
		double x[] = g.rezolvare(100);
		
		for(int i=0; i<x.length; i++)
			System.out.print(x[i]+" ");
	}
}