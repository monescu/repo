import java.io.File;
	import java.io.FileNotFoundException;
	import java.util.Scanner;
	
	public class Iterativa{
	
		static int n;
		static double[] x, d, a, b, r, c, s;
	
		static double[] citireSir(Scanner in) {
	
			String[] rand = in.nextLine().split(" ");
			double[] sir = new double[rand.length + 1];
			for (int i = 1; i <= rand.length; i++) {
				sir[i] = Double.parseDouble(rand[i - 1]);
			}
			return sir;
		}
	
		public static void main(String[] args) throws FileNotFoundException {
	
			File file = new File("fisier.in");
	
			Scanner sc = new Scanner(file);
			a = citireSir(sc);
			n = a.length - 1;
	
			System.out.print("a = ");
			for (int i = 1; i <= n; i++)
				System.out.print(a[i] + " ");
			System.out.println();
	
			b = citireSir(sc);
	
			System.out.print("b = ");
			for (int i = 1; i <= n - 1; i++)
				System.out.print(b[i] + " ");
			System.out.println();
	
			c = citireSir(sc);
	
			System.out.print("c = ");
			for (int i = 1; i <= n - 1; i++)
				System.out.print(c[i] + " ");
			System.out.println();
			d = citireSir(sc);
	
			System.out.print("d = ");
			for (int i = 1; i <= n; i++)
				System.out.print(d[i] + " ");
			System.out.println();
	
			r = new double[n];
			x = new double[n + 1];
			s = new double[n];
	
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
	
			for (int i = 1; i <= n; i++)
				System.out.println(x[i]);
		}
	
	}


