// Compile with
// g++ --std=c++11 cpp.cpp -o cpp.out
#include <iostream>

int main() {
    /* works: bool trues = false;*/
    true = false;  // Does not work:  error: lvalue required as left operand of assignment
    if (true) {
        std::cout << "C++11 does not 'support' reassignment of true :-)" << std::endl;
    } else {
        std::cout << "C++11 'supports' reassignment of True :-(" << std::endl;
    }

}