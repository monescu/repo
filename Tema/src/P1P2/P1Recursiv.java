package P1P2;
	import java.util.Scanner;
	import java.io.File;

	public class P1Recursiv {

		static float[] x, d, a, b, r, c, s;
		static int n;

		public static void rezolvareRsiS(float[] r, float[] s, int i) {
			if (i == 1 && a[1] != 0) {
				r[1] = -c[1] / a[1];
				s[1] = d[1] / a[1];
			}
			if (i < n - 1) {
				r[i + 1] = -c[i + 1] / (b[i] * r[i] + a[i + 1]);
				s[i + 1] = (d[i + 1] - b[i] * s[i]) / (b[i] * r[i] + a[i + 1]);
				rezolvareRsiS(r, s, i + 1);
			}
		}

		public static void rezolvareX(float[] r, float[] s, int i) {
			if (i == n) {
				x[i] = (d[i] - b[i - 1] * s[i - 1]) / (b[i - 1] * r[i - 1] + a[i]);
				rezolvareX(r, s, i - 1);
			} else if (i > 0) {
				x[i] = r[i] * x[i + 1] + s[i];
				rezolvareX(r, s, i - 1);
			}
		}

		static float[] input(Scanner sc) {

			String line[] = sc.nextLine().split(" ");
			float vect[] = new float[line.length + 1];
			for (int i = 1; i <= line.length; i++) {
				vect[i] = Float.valueOf(line[i - 1]);
			}
			return vect;
		}

		public static void main(String[] args) throws Exception {

			File dateProblema = new File("P1.in");
			Scanner sc = new Scanner(dateProblema);
			a = input(sc);
			n = a.length - 1;
			System.out.println("Matricea este:");
			System.out.print("a={ ");
			for (int i = 1; i <= n; i++)
				System.out.print(a[i] + " ");
			System.out.println("}");
			b = input(sc);
			System.out.print("b={ ");
			for (int i = 1; i <= n - 1; i++)
				System.out.print(b[i] + " ");
			System.out.println("}");
			c = input(sc);
			System.out.print("c={ ");
			for (int i = 1; i <= n - 1; i++)
				System.out.print(c[i] + " ");
			System.out.println("}");
			d = input(sc);
			System.out.print("d={ ");
			for (int i = 1; i <= n; i++)
				System.out.print(d[i] + " ");
			System.out.println("}");

			r = new float[n];
			x = new float[n + 1];
			s = new float[n];
			
			rezolvareRsiS(r, s, 1);
			rezolvareX(r, s, n);
			
			System.out.println("\nSolutii:");
			for (int i = 1; i <= n; i++)
				System.out.println("x[" + i + "]: " + (float) Math.round(x[i] * 100000d) / 100000d);
		}

	}
