package primaProblema;

import java.io.File;
import java.util.Scanner;

public class Iterativ {

	static float[] citireVectorFisier(Scanner sc) {

		String line[] = sc.nextLine().split(" ");
		float vect[] = new float[line.length + 1];
		for (int i = 1; i <= line.length; i++) {
			vect[i] = Float.valueOf(line[i - 1]);
		}
		return vect;
	}

	public static void main(String[] args) throws Exception {

		int n;
		float[] x, d, a, b, r, c, s;

		File txt = new File("Matrice.txt");
		Scanner sc = new Scanner(txt);
		a = citireVectorFisier(sc);
		n = a.length - 1;
		System.out.println("Matricea este:");
		System.out.print("a: ");
		for (int i = 1; i <= n; i++)
			System.out.print(a[i] + " ");
		System.out.println();
		b = citireVectorFisier(sc);
		System.out.print("b: ");
		for (int i = 1; i <= n - 1; i++)
			System.out.print(b[i] + " ");
		System.out.println();
		c = citireVectorFisier(sc);
		System.out.print("c: ");
		for (int i = 1; i <= n - 1; i++)
			System.out.print(c[i] + " ");
		System.out.println();
		d = citireVectorFisier(sc);
		System.out.print("d: ");
		for (int i = 1; i <= n; i++)
			System.out.print(d[i] + " ");
		System.out.println();

		r = new float[n];
		x = new float[n + 1];
		s = new float[n];

		if (a[1] != 0) {
			r[1] = -c[1] / a[1];
			s[1] = d[1] / a[1];
		}

		for (int i = 1; i <= n - 2; i++) {
			r[i + 1] = -c[i + 1] / (b[i] * r[i] + a[i + 1]);
			s[i + 1] = (d[i + 1] - b[i] * s[i]) / (b[i] * r[i] + a[i + 1]);
		}

		x[n] = (d[n] - b[n - 1] * s[n - 1]) / (b[n - 1] * r[n - 1] + a[n]);

		for (int i = n - 1; i >= 1; i--)
			x[i] = r[i] * x[i + 1] + s[i];
		System.out.println("\nSolutile:");
		for (int i = 1; i <= n; i++)
			System.out.println("x[" + i + "]: " + (float) Math.round(x[i] * 100000d) / 100000d);
	}

}
