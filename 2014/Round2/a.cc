#include <iostream>
#include <algorithm>
#include <vector>

int solve()
{
	unsigned long int N, X;
	std::vector<unsigned long int> files;

	std::cin >> N >> X;

	for (unsigned long int i = 0; i < N; i++) {
		unsigned long int file;
		std::cin >> file;
		files.push_back(file);
	}

	std::sort(files.begin(), files.end());

	for (unsigned long int i = 0, j = files.size() - 1; i < j; i++) {
		while (i < j && files[i] + files[j] > X) {
			j--;
		}
		if (i < j) {
			N--;
			j--;
		}
	}
	return N;
}

int main(int argc, char *argv[])
{
	unsigned long int ncases;

	std::cin >> ncases;

	for (unsigned long int i = 1; i <= ncases; i++) {
		std::cout << "Case #" << i << ": " << solve() << std::endl;
	}

	return 0;
}