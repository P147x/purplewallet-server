#include <string>
#include <iostream>
#include <termios.h>
#include <unistd.h>
#include "config.hpp"
#include "network.hpp"

#ifndef PURPLE_HPP
#define PURPLE_HPP

class Purple {
    private:
        Config      config;
        Network     network;
        bool        login();
    public:
        bool        tryLogin();
        int         run();
        Purple();

    int run(std::vector<std::string> args);

    void commandPicker(char **argv);

    void commandPicker(std::vector<std::string> args);

    void getWalletInformation(int id);

    void logout();

    void addPurchase(bool isDebt);

    float setSum();
};

#endif //PURPLE_HPP