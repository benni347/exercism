#include "reverse_string.h"
#include <algorithm>

namespace reverse_string {

std::string reverse_string(const std::string &input) {
  std::string reversed(input);
  std::reverse(reversed.begin(), reversed.end());
  return reversed;
}

} // namespace reverse_string
