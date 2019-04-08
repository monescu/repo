import java.io.*;
//import java.util.Arrays;
import java.util.StringTokenizer;

public class Tdma {
	 
	public static double[] solve(double[] a, double[] b, double[] c, double[] d) 
	{ 
	  int n = d.length; 
	  double temp; 
	  c[0] /= b[0]; 
	  d[0] /= b[0]; 
	  for (int i = 1; i < n; i++) { 
	   temp = 1.0 / (b[i] - c[i - 1] * a[i]); 
	   c[i] *= temp; // redundant at the last step as c[n-1]=0. 
	   d[i] = (d[i] - d[i - 1] * a[i]) * temp; 
	  } 
	  double[] x = new double[n]; 
	  x[n - 1] = d[n - 1]; 
	  for (int i = n - 2; i >= 0; i--) { 
	   x[i] = d[i] - c[i] * x[i + 1]; 
	  } 
	  return x; 
	} 
	 
	public static void main(String[] args) throws IOException
	{
	    int n;
	    double[][] M;

	    BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));
	    PrintWriter writer = new PrintWriter(System.out, true);

	    n = Integer.parseInt(reader.readLine());
	    M = new double[n][n+1];
	    
	    for (int i = 0; i < n; i++) {
	        StringTokenizer strtk = new StringTokenizer(reader.readLine());

	        while (strtk.hasMoreTokens())
	        for (int j = 0; j < n + 1 && strtk.hasMoreTokens(); j++)
	          M[i][j] = Integer.parseInt(strtk.nextToken());
	      }

		double[] a = new double[] {5,4,6,4};
		double[] b = new double[] {5,4,6,4};
		double[] c = new double[] {5,4,6,4};
		double[] d = new double[] {0,4,6,4};
		double[] result = Tdma.solve(a, b, c, d);
		writer.println(result);
	}

}
