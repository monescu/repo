package P1P2;
import java.io.File;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.util.Scanner;

public class P2GaussSeitel {
	private double[][] A;
	private double[] b,x,y;
	private int n;
	private double eps;
	private long nrMax;
	private double calcNorm() {
		double max = Math.abs(x[0] - y[0]);
	for(int i=1;i<n;i++) {
		double calc = Math.abs(x[i] - y[i]);
		if(calc > max)
			max = calc;
	}
	return max;
}
private double calcSum(int row) {
	double sum = 0.0;
	
	for(int i = 0;i<n;i++)
		if(i != row)
			sum += A[row][i] * x[i];
	return sum;
}

private double calcSumM(int row) {
	double sum = 0.0;
	for(int i = row;i < n;i++)
		if(i != row)
			sum += A[row][i] * x[i];	
	return sum;
}
private double calcSumU(int row) {
	double sum = 0.0;
	for(int i = 0;i<row;i++)
		if(i != row)
			sum += A[row][i] * y[i];
	return sum;
}
public P2GaussSeitel(String path,long nrMax,double eps) {
	try {
		this.eps = eps;
		this.nrMax = nrMax;
		Scanner sc = new Scanner(new FileReader(new File(path)));
		n = sc.nextInt();
		A = new double[n][n];
		b = new double[n];
		x = new double[n];
		y = new double[n];
		double norm;
		for(int i =0 ;i<n;i++) 
			for(int j =0 ;j<n;j++) 
				A[i][j] = Double.parseDouble(sc.next());
		for(int i =0 ;i<n;i++)
				b[i] = Double.parseDouble(sc.next());
		sc.close();
		int nrIt = 0;
		do {
			//y[0] = (1.00/A[0][0]) * (b[0] - calcSum(0));
			for(int i=0;i<n;i++) 
				y[i] = (1.00/A[i][i]) * (b[i] - calcSumU(i) - calcSumM(i));
			norm  = calcNorm();
			for(int i=0;i<n;i++)
				x[i] = y[i];
			nrIt++;
		}while(norm >= eps && nrIt < nrMax);
		if(norm < eps)
			System.out.println("Precis.");
		else
			System.out.println("Nu precis.");
		for(int i=0;i<n;i++)
			System.out.println(y[i]);
	} catch (FileNotFoundException e) {
		e.printStackTrace();
	}
}
	public static void main(String args[]) {
	P2GaussSeitel gs = new P2GaussSeitel("P2.in", 10000, 0.0001);
}
}