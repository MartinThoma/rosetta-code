#include <stdio.h>      // printf
#include <vector>
#include <random>       // random number generation
#include <algorithm>    // sorting algorithm
#include <chrono>       // benchmarking
using namespace std;

int main() {
    typedef std::chrono::high_resolution_clock Clock;
    typedef std::chrono::duration<double> sec;

    int n = 10000000;
    int min = 0;
    int max = 1000000000;

    vector<int> myVector;  // create an empty vector

    /* initialize random seed: */
    srand (time(NULL));

    for (int i = 0; i < n; ++i) {
        int output = min + (rand() % (int)(max - min + 1));
        myVector.push_back(output);
    }

    Clock::time_point t0 = Clock::now();
    std::sort (myVector.begin(), myVector.end());
    Clock::time_point t1 = Clock::now();

    printf("%0.2f seconds\n", sec(t1-t0).count());
    return 0;
}