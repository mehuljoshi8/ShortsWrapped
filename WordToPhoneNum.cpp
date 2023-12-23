// Author: Mehul Joshi
// Simple program to convert a word to a phone number
// by default we require that the word only be 10 characters long at max.

#include <iostream>
#include <vector>
#include <stack>
#include <string>
#include <sstream>

using std::cout;
using std::cin;
using std::endl;
using std::vector;
using std::string;
using std::stringstream;
using std::stack;

vector<int> charToDig {2,2,2,3,3,3,4,4,4,5,5,5,6,6,6,7,7,7,7,8,8,8,9,9,9,9};

void letterCombosHelper(string digits, int i, stack<char> curr, vector<string>& res, vector<vector<char>> mapping) {
    if (i == digits.size()) {
        stringstream ss;
        stack<char> tmp = curr;
        while (!tmp.empty()) {
            ss << tmp.top();
            tmp.pop();
        }
        string to_add = ss.str();
        reverse(to_add.begin(), to_add.end());
        res.push_back(to_add);
    } else {
        for (int k = 0; k < mapping[digits[i] - '0' - 2].size(); k++) {
            curr.push(mapping[digits[i] - '0' - 2][k]);
            letterCombosHelper(digits, i+1, curr, res, mapping);
            curr.pop();
        }
    }
}

vector<string> letterCombinations(string digits) {
    vector<string> res;
    if (digits.size() == 0)
        return res;
    stack<char> curr;
    vector<vector<char>> mapping = {{'a', 'b', 'c'},
                                    {'d', 'e', 'f'},
                                    {'g', 'h', 'i'},
                                    {'j', 'k', 'l'},
                                    {'m', 'n', 'o'},
                                    {'p', 'q', 'r', 's'},
                                    {'t', 'u', 'v'},
                                    {'w', 'x', 'y', 'z'}};
    letterCombosHelper(digits, 0, curr, res, mapping);
    return res;
}

int main() {
    vector<string> letter_combos = letterCombinations("6954737");
    for (string combo: letter_combos) {
        cout << combo << endl;
    }
    // letter_combos = letterCombinations("6954737");
    // for (string combo: letter_combos) {
    //     cout << combo << endl;
    // }
    return 0;
}
