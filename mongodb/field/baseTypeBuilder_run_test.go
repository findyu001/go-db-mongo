package field

import (
  "fmt"
)

/**
以下两个命令，是为了生成库文件，上层使用方无需执行下列命令，但即使执行了，也不会也任何改变
需要先执行 Example_buildAll0 再执行 Example_buildAll1
 */

func Example_buildAll0() {
  buildAll0()

  fmt.Println(true)
  // Output:
  // true
}

func Example_buildAll1() {
  buildAll1()

  fmt.Println(true)
  // Output:
  // true
}
