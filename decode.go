package main

import (
    "crypto/md5"
    "io"
    "fmt"
    "time"
    "encoding/hex"
    "os"
)

const (
    STRING = "0123456789abcdefghijklmnopqrstuvwxyz"
    CODE   = "eae8a7f639d78f71f55ef100bf7f3e45"
)

func reverseStrings(s string) string {
    str := []rune(s)

    for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
        str[i], str[j] = str[j], str[i]
    }

    return string(str)
}

func main() {
    reverseString := reverseStrings(STRING)
    fmt.Println(time.Now())

    // 一位密码
    go func() {
        for _, one := range STRING {
            h := md5.New()
            io.WriteString(h, string(one))
            str := hex.EncodeToString(h.Sum(nil))

            if CODE == str {
                fmt.Println("密码已经找到,一位密码:", string(one))
                os.Exit(0)
            }
        }
        fmt.Println("不是一位密码!", time.Now())
    }()

    // 两位密码
    go func() {
        for _, one := range STRING {
            for _, two := range STRING {
                h := md5.New()
                io.WriteString(h, string(one)+string(two))
                str := hex.EncodeToString(h.Sum(nil))
                if CODE == str {
                    fmt.Println("密码已经找到，两位密码:", string(one)+string(two))
                    os.Exit(0)

                }
            }
        }
        fmt.Println("不是两位密码!", time.Now())
    }()

    // 三位密码
    go func() {
        for _, one := range STRING {
            for _, two := range STRING {
                for _, three := range STRING {
                    h := md5.New()
                    io.WriteString(h, string(one)+string(two)+string(three))
                    str := hex.EncodeToString(h.Sum(nil))

                    if CODE == str {
                        fmt.Println("密码已经找到，三位密码:", string(one)+string(two)+string(three))
                        os.Exit(0)
                    }
                }
            }
        }
        fmt.Println("不是三位密码!", time.Now())
    }()

    // 四位密码
    go func() {
        for _, one := range STRING {
            for _, two := range STRING {
                for _, three := range STRING {
                    for _, four := range STRING {
                        h := md5.New()
                        io.WriteString(h, string(one)+string(two)+string(three)+string(four))
                        str := hex.EncodeToString(h.Sum(nil))

                        if CODE == str {
                            fmt.Println("密码已经找到，四位密码:", string(one)+string(two)+string(three)+string(four))
                            os.Exit(0)
                        }
                    }
                }
            }
        }
        fmt.Println("不是四位密码!", time.Now())
    }()

    //五位密码
    go func() {
        for _, one := range STRING {
            if string(one) == "d" {
                break
            }
            for _, two := range STRING {
                for _, three := range STRING {
                    for _, four := range STRING {
                        for _, five := range STRING {
                            h := md5.New()
                            io.WriteString(h, string(one)+string(two)+string(three)+string(four)+string(five))
                            str := hex.EncodeToString(h.Sum(nil))

                            if CODE == str {
                                fmt.Println("密码已经找到，五位密码:", string(one)+string(two)+string(three)+string(four)+string(five))
                                os.Exit(0)
                            }
                        }
                    }
                }
            }
        }
        fmt.Println("不是五位密码(0-c)！", time.Now())
    }()

    // 五位逆序查找
    go func() {
        for _, one := range reverseString {
            if string(one) == "c" {
                break
            }
            for _, two := range reverseString {
                for _, three := range reverseString {
                    for _, four := range reverseString {
                        for _, five := range reverseString {
                            h := md5.New()
                            io.WriteString(h, string(one)+string(two)+string(three)+string(four)+string(five))
                            str := hex.EncodeToString(h.Sum(nil))

                            if CODE == str {
                                fmt.Println("密码已经找到，五位密码:", string(one)+string(two)+string(three)+string(four)+string(five))
                                os.Exit(0)
                            }
                        }
                    }
                }
            }
        }
        fmt.Println("不是五位密码(d-z)！", time.Now())
    }()

    recode()
    time.Sleep(600 * time.Second)
}

func recode() {
    // 全部1-6位数查找时间，19:54:58----19:55:30, 数据：s1a3z5
    for _, one := range STRING {
        ones := string(one)
        go func() {
            for _, two := range STRING {
                for _, three := range STRING {
                    for _, four := range STRING {
                        for _, five := range STRING {
                            for _, six := range STRING {
                                h := md5.New()
                                io.WriteString(h, ones+string(two)+string(three)+string(four)+string(five)+string(six))
                                str := hex.EncodeToString(h.Sum(nil))

                                if CODE == str {
                                    fmt.Println("密码已经找到，六位密码:", ones+string(two)+string(three)+string(four)+string(five)+string(six), time.Now())
                                    os.Exit(0)
                                }
                            }
                        }
                    }
                }
            }
            fmt.Println("六位密码首字母：", ones, "未找到密码！", time.Now())
        }()
    }
}
