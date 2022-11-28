#include <iostream>
#include <omp.h>
#include <thread>
#include <chrono>

using namespace std;
void open_mp_check()
{
#ifdef _OPENMP
    printf("Open Mp enabled");
#else
    printf("open Mp disabled");
#endif
}
double result = 1;
void wallisP() {
    #pragma omp parallel for
    for (int i = 1; i < 1000; i++)
    {
        std::this_thread::sleep_for(std::chrono::nanoseconds(100));
        double num = 4.0 * i * i;
        result *= num / (num - 1);
    }
    result *= 2;
    
}

int main()
{
    open_mp_check();
    cout << endl;
    double begin = 0, end = 0;
    
    begin = omp_get_wtime();
    wallisP();
    end = omp_get_wtime();

    cout << "Wallis parallel implementation: " << end - begin << endl;

    return 0;
}
