#include <stdlib.h>

int points(const char* const games[10]) {
  int total_x = 0;

  
  for(int i=0; i < 10; i++){
    int x = games[i][0] - '0';
    int y = games[i][2] - '0';
    
    if(x > y){
      total_x += 3;
    }
    
    else if(x == y){
      total_x += 1;
    }
    
    else {
      continue;
    }
  }
  
  return total_x;
}
