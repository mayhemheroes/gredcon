package fuzz

import "strconv"
import "github.com/nisainan/gredcon/gredcon"

func mayhemit(bytes []byte) int {

    var num int
    if len(bytes) < 1 {
        num = 0
    } else {
        num, _ = strconv.Atoi(string(bytes[0]))
    }
    switch num {
        case 1:
            gredcon.ReadCommands(bytes)
            return 0

        case 2:
            gredcon.ReadNextRESP(bytes)
            return 0

        case 3:
            content := string(bytes)
            gredcon.AppendString(bytes, content)
            return 0
        case 4:
            gredcon.AppendOK(bytes)
            return 0
        default:
            var test gredcon.Handler
            var test2 gredcon.ServeMux
            content := string(bytes)
            test2.Handle(content, test)
            return 0
    }
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}