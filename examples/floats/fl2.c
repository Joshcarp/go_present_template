#include <stdio.h>
 
int main(void) {
  long double a = 1.0; 
  long double b = 3.0; 
  double f = (a / b) - (1.0/3.0);
  printf("%.50f\n", f);
  printf("%ld\n", sizeof(long double));
}