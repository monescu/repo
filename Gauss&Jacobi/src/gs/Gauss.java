package gs;
public class Gauss {
		static double a[][] = { {15, 1, 7 }, {17, 33, -1}, {-2, 16, 74} };
		static double b[] = {6, 7, 8};
		static double norm=0;
		static int n=b.length;
		static final long nrMax=1000000;
		static double x[], y[];
		static double epsi= 0.0001;
		static long nrIt=0;

		public static double norm() {
			double max=Math.abs(x[0] - y[0]);
			for (int i=1; i<n; i++) {
				double diff=Math.abs(x[i]-y[i]);
				if (diff>max)
					max=diff;
			}
			return max;
		}

		public static double sum1(int l) {
			double sum=0;
			for (int i=0; i<l; i++) {
				if (l!=i)
					sum+=a[l][i]*y[i];
			}
			return sum;
		}

		public static double sum2(int l) {
			double sum=0;
			for (int i=l; i<n; i++) {
				if (l!=i)
					sum+=a[l][i]*x[i];
			}
			return sum;
		}

		public static void rez() {
			x=new double[100];
			y=new double[100];
			do {
				for (int i=0; i<n; i++) {
					y[i]=(1.00/a[i][i])*(b[i]-sum1(i)-sum2(i));
				}
				norm=norm();
				nrIt++;
				for (int i=0; i<n; i++)
					x[i]=y[i];
			} while (norm>epsi && nrIt<nrMax);
		}

		public static void afis() {
			if (norm<epsi)
				System.out.println("Ok.");
			else
				System.out.println("Not ok.");

			for (int i=0; i<n; i++)
				System.out.println("x["+i+"]: "+y[i]);
		}

		public static void main(String args[]) {
			System.out.println("Matricea:");
			System.out.print("a: ");
			for (int i=0; i<a.length; i++) {
				for (int j=0; j<a.length; j++)
					System.out.print(a[i][j]+" ");
				if (i<a.length - 1)
					System.out.print("\n   ");
			}
			System.out.print("\n b: ");
			for (int i=0; i<b.length; i++)
				System.out.print(b[i]+" ");
			System.out.println("\n Solutii:");
			rez();
			afis();
		}
	}

