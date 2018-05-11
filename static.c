#include <stdio.h>

void numberOfCalls() {

    // Initialized to 0 at compile time
    static int counter = 0;

    counter++;

    printf("This function has been called %d times\n", counter);
}


int main() {

    for (int i = 0; i < 5; i++) {
        numberOfCalls();
    }

    return 0;
}