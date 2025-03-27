#include <iostream>
#include <fstream>

int main(int argc, char* argv[]) {
    if (argc < 3) {
        std::cerr << "invalid arguments" << std::endl;
        return 1;
    }

    std::string manifest_path = argv[1];
    std::string file_path = argv[2];

    if (!std::filesystem::exists(manifest_path)) {
        std::cerr << "manifest file does not exist" << manifest_path << std::endl;
        return 1;
    }

    if (!std::filesystem::exists(file_path)) {
        std::cerr << "manifest file does not exist" << manifest_path << std::endl;
        return 1;
    }
}