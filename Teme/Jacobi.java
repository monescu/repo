import java.util.Scanner;
public class Jacobi {

	public static double max(double a, double b){
		if (a>b)
			return a;
		return b;
	}

	public static double norma(double[] x, double[] y, int n){
		double max=0;
		for (int i=0; i<n; i++) {
			max=max(max, Math.abs(x[i]-y[i]));
		}
		
		return max;
	}

	public static void main(String[] args){
		int n;
		Scanner sc=new Scanner(System.in);
		System.out.println("Lungime:");
		n=sc.nextInt();
		double a[][]=new double[n][n];
		int k=1;
		int kmax=500;
		System.out.println("Matricea: ");
		for (int i=0; i<n; i++)
			for (int j=0; j<n; j++) {
				a[i][j]=sc.nextDouble();
			}
		double y[]=new double[n];
		double x[]=new double[n];
		for (int i=0; i<n; i++)
			x[i]=0;
		double b[] = new double[n];
		System.out.println("b=");
		for (int i=0; i<n; i++)
			b[i]=sc.nextDouble();
		double epsilon=0.001;
		double norma;
		do {
			for (int i=0; i<n; i++) {
				double c=0;
				for (int j=0; j<n; j++) {
					if (j!=i) {
						c=c+a[i][j]*x[j];
					}
				}
				y[i]=(1/a[i][i])*(b[i]-c);
			}
			k++;
			norma=norma(x, y, n);
			for (int i=0; i<n; i++) {
				x[i]=y[i];
			}

		} while ((norma>=epsilon) && (k<kmax));
		System.out.println("Solutia: ");
		for (int i=0; i<n; i++)
			System.out.println(x[i] + " ");
	}

}
