#include <stdbool.h>

bool better_than_average(const int class_points[], int class_size, int your_points) {
  double total = 0;
  double avg = 0;
  for(int i=0; i < class_size; i++)
    {
    total += class_points[i];
  }
  
  avg = total / class_size;
  
  if (your_points > avg) {
    return true;
  }
  
  else {
    return false;
  }
}
