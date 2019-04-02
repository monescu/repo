package aDouaTema;

//METODA GAUSS-SEIDEL

public class Main {

	public static void main(String[] args) {
		double[][] A = { { 16, 3 }, { 7, -11 } };
		double[] b = { 11, 13 };
		MetodaGaussSeidel solver = new MetodaGaussSeidel(A, b);
		double[] x = solver.solveSystem(100);
		for (int i = 0; i < x.length; i++)
			System.out.println("Prima solutie: " + x[i]);

	}

}
