package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "5.3" )

func h0tDy7kJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AuyUYtSBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vPlJx9AMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SSiBNbNdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lw462mDdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VLRr0omXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tgmez5QbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jc74mNyLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mxlnpit3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SJCUMyAxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qc2nMfnhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zO1EDVarWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nvVVpobeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gCr09hZQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xz0EMVPuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tz2ydEa2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eP5PYYshWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OPafCNyeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PzCtY9D2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F5GFFTDeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Hak1LadWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SFDV8lGfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FCfc6p43Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V4D16EkEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6sP5TdnCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gZwsfjAIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xKq8AOBZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XAlx3zr5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R3Xixl0gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nd7lQhtgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rZbwA4MYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4nNAVYSWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rvI0b9plWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h2WIl0i2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N4FsParyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tNIaL2N1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B5jt92JLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1xSsnomKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ecJwJvNDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t7o1qKEYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YIiETz32Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QasEsDPaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tXSlD20JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZFNxrQGTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aPvDUcOlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YZNq3EKyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bGOyEGhCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xx1fZ3sSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func StxBblDaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func luzNvQBJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 91YErp3FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ATuuA9LIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A7FVufhfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vwnS9EZEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aAJBv5qxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eQY4SmokWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sHYDDBw5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8k9rFHhQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kWeEhjBqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KNXzXAITWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kQonQXx6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wW5NcRfHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vDIq7P8PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IRpkCKdPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q5lBLANaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func otHqjuZ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UoMO6lrFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func giONxHx0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9VmiGyNJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func toiDY9GZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SiUfFSfyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vyzc2O7WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rYNEXMQVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RpvZljtGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0iaoUdIAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tgfpiSb6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rDb1kjLuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jJ9dB95jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RenDMPfWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cv6GUJHRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J8I2T61YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I4NZFMs0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KwJgThCMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TYtuJrqTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tJQJyF3aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OsytTru8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NueNHb0AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VSORdOicWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wh4ZePNUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6EGmfaHkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 575OKh6HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Znh66Be1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vFpdmDtDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9HK7jeQiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ywY3qZhTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wS3vfOGzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ClV2sxkmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NZB89KHeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w2kwwDIfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x9P1XweAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8tTVKQhaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aeBwIyRvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EfWVXZ2LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p0z1KyYmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qGwAItEbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JaXl9mjVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lu3iyJb0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3ivaz42PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dE4Byk6oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c9RXssDVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0RsuFpFHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UkWTJuQlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RjUw3u3eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6GFvfTJbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A6VuE0IYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2O86wbPuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KgcWiDgqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7G6hcitNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dgNfbHBNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jof0Q5guWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8TMrE6jzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Tbo3xjXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kFOW8dVCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hmiptEc5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5b3uKbGBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2YIZdb4pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ROCIuJyrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W5OF9iXwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I47PvDOcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s9nqB3TUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QoWWOwSuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pi9tjQZFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oDvi0AARWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GA8yXV1tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5mkgT8k7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3uLxEpP7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tq8uoFb9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NT1Uyir2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0dkf2x36Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aWa4Yoh7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LhYSK1NKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gXx1c7QCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dAyZoB9QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0xvCAOf5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 80x8EZM1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O48Y9gM4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KXg3zi8MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WsaIJqZdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nQKlcFDwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r6UUYl8hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LnXcT2nxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xC2drUcNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S7k7lkJMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vpgpIIB2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H9de9P4HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y4fms32ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j1w1qJieWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RAkCaOQNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nLdgdvFqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func woNnSexZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ciXJtH9WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vrc2sE2SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SGdesGpsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2uG85mMCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 838f1AhzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PPVUd2ghWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 58X0a98CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5UkslYVLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bmAbhQgvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZkMKpmyLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ltTFivXIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2xDWZH1LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4bV0VlpCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hIWZ1z6cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9CqHs6XBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dD99SdmNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PLWyO6FhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PKBzZQL6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g26j81kWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wVZFwVQZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tXikXFkbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s2Ie4PSTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TZAHUU7AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RUMzwYPwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bKgoGA1wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k2OcIrMcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jjsntPPnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mgRgxqksWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LBhEYBqCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D6IfV9DNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zT6HdnMTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dyHfASwpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tRl2kQ8bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BmJpd68wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NdEDdpolWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K71WBbDEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kZrXZf2CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TAz3MHZGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ciOAjKGDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aGYj4FDUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wVwaE4UJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nDTALEHsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uDamoHbnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5U8hEH3EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o2Jw1YEmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eJkmYtfYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uD2JPWdDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JfwJVEbCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YcnCX4ShWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O5usUK4HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JFwkSaChWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nB1r64GKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9d25XcZ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xqYV8wohWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uynKaW4DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Np1iNortWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lcUg28b7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i5so5lx5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IGZ7iv8wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kyigiOIkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r2DNL5HaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zv3W0SSxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QXaImYGMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ymJWBDcGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8ctD23ogWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3MEoEGPWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rOtuGxFnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func orDkto6qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7TrPrtjbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2n3L3iUCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fV1WZaXzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JqgvG2MzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NsBx0HprWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vIDoK9CdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func udZg9LDIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fSSqPGVkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UaIHslDkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func noMbTuSPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I5zPa2QkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lLeO99QNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T1gdi694Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6eEcwfQ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rpYcWJVmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A2rCT9SSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q6AGYn42Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YbKWW4u1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q1mQ3s2tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZOGeIRRKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sgnWOmZEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VkaUmgxKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fYFKMwcwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AfbLx3nkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jE5nUY7ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DFo0hjteWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hd1gZSR9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B6iIT35jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2TKsdkscWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VAcnajzKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rkjV60B1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gaZ6egBTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qtbTqZV7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6nL5izG1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lcwVaXiDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V0DSR4CtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AaBcKqmzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 670Cw3vzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YfOn8Tb6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5j8SddGjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wWOw6R6DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func foZ7uRQ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Glo7OdC9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l3iFUlYYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YZFKUYRYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S7xCcp57Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E4ueHAUMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xEIJvbXIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rm7vySBjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AC9FZ4k7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xiCqdHGjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 01WIJAwgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9CzNS5QiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RuMEubAxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NLpwmgVlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ikHK4EaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sy44rWYUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9SPZgna7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jlTM0vsVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e7vb51hMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bGFZ2j6OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K6gCpb02Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xKdHTQjyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VpwP1ZPtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VTlVivb8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GWORXKLNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ufCSSDJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rSdQcc5QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dhk8ECBDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OjN8M4GaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6NbN7lN0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9vuO9uwhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gYhqU770Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Vs2arCPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GUc4UcS1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tp4nmyzsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Y7hZARrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2XRHokpmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BBeLHVYnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PIKs6tnwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ZobighiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7lAhTJiaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gdGZL0ENWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QqPzewfZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QqftRSdaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TZIXeFbZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BqrMHxUJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jpP1TlnEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u0WinBngWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fmWHEfFzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IInagI0YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c2HS4HZ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C9PlSn00Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uYhIT17dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OFMQWrLjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zYSbPtQzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dvJvd3obWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eivid6UHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Khg0NaEfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GTcSknmBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rf4LtmPPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func btoHBYJmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RGseEMagWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sTQTtkbIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hRQaoL1xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rIbJVlhwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EJpN92A9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func St9Wfv2DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BW2IxzzLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AiMuE7p8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g7TanS4FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func auEqVu5tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UNlultfIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2KAoX8edWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BQxi9pPqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tSA78QJhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0qNt7G4yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6v59SXxGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VlOilABEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sJYqfaqOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PgaQPMZ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JOOgL8HkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZpkCPZY9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lAXvtBSvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ypPy78H1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9nCy7JI8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jlj340ueWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dsJEzvV3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zA7iNZ70Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W58EoXnrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u4iGXoGGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ruKimr5aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f0hWtrERWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OXeJkqzOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0zt1BbU0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DZiLoc9sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cfJ3Y2ZuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hg9hMvDgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZyiH6A8GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oxXQ1iPFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zcHfNa0RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6jxq2irSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GZqSm63SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PECXSfgfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8bdgjrWTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uAsuglUwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func odgAEgbNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pTqbpkeJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pXeU1ks0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 637KKWvgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cw5NarmKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yV43zFnfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oTnqoNRUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kTIn603rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bPni4sN6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

