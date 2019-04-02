import java.util.Scanner;
public class Recursiv {

	static float[] x, d, a, b, r, c, s;
	static int n;

	public static void rS(float[] r, float[] s, int i){
		if (i==1&&a[1]!=0) {
			r[1]=-c[1]/a[1];
			s[1]=d[1]/a[1];
		}
		if (i<n-1){
			r[i+1]=-c[i+1]/(b[i]*r[i]+a[i+1]);
			s[i+1]=(d[i+1]-b[i]*s[i])/(b[i]*r[i]+a[i+1]);
			rS(r, s, i+1);
		}
	}

	public static void xA(float[] r, float[] s, int i) {
		if (i==n) {
			x[i]=(d[i]-b[i-1]*s[i-1])/(b[i-1]*r[i-1]+a[i]);
			xA(r, s, i-1);
		} else if (i>0) {
			x[i]=r[i]*x[i+1]+s[i];
			xA(r, s, i-1);
		}
	}

	static float[] citire(Scanner sc) {

		String line[] = sc.nextLine().split(" ");
		float vect[] = new float[line.length + 1];
		for (int i = 1; i <= line.length; i++) {
			vect[i] = Float.valueOf(line[i - 1]);
		}
		return vect;
	}

	public static void main(String[] args){

		Scanner sc=new Scanner(System.in);
		System.out.println("Matricea:");
		a=citire(sc);
		n=a.length - 1;
		System.out.print("a: ");
		for (int i=1; i<=n; i++)
			System.out.print(a[i]+" ");
		System.out.println();
		b=citire(sc);
		System.out.print("b: ");
		for (int i=1; i<=n-1; i++)
			System.out.print(b[i]+" ");
		System.out.println();
		c=citire(sc);
		System.out.print("c: ");
		for (int i=1; i<=n-1; i++)
			System.out.print(c[i]+" ");
		System.out.println();
		d=citire(sc);
		System.out.print("d: ");
		for (int i=1; i<=n; i++)
			System.out.print(d[i]+" ");
		System.out.println();

		r=new float[n];
		x=new float[n+1];
		s=new float[n];

		rS(r, s, 1);
		xA(r, s, n);

		System.out.println("\n Solutii:");
		for (int i=1; i<=n; i++)
			System.out.println("x["+i+"]:"+x[i]);
	}

}
