package alipay_test

import (
	"testing"

	"github.com/crazytaxii/alipay"
)

const (
	APP_ID    = "APP_ID"
	SELLER_ID = "PID"

	PEM_MY_PUB_KEY = `
-----BEGIN PUBLIC KEY-----
MIIBITANBgkqhkiG9w0BAQEFAAOCAQ4AMIIBCQKCAQBaoXIExBjF7bS20YAA5jJD
5E995eDg1lPn7lUB7TaEKHeiZ4hFZIdowLp6vngIAiR/BPAbLANp6FJ0Y7VrDTdO
otJwU6vcW2Q0NVdFOsQp1GdaFU48+x+h9VP442jDk36wz/xPQw7DvrJ/zfcr+x0m
n2vkVShrf4V41RnIh7ATj+SGr764ilpT4np12S1MLDfPZaAM5OcIoaQQdH6l8C6O
8FNoS1WEE0yk3s6RG54Yodjk2cxb9NaOBAbqiKwcg+TRPoBLDcFQrI8gEeq+8m6G
X3Q5bmC6IOZh+LzvTkqYuoQKWjJy+cTFZwlEYrqe5ab+XLGEVnZSr6Xdplt4QGYV
AgMBAAE=
-----END PUBLIC KEY-----
`

	PEM_MY_PVT_KEY = `
-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQBaoXIExBjF7bS20YAA5jJD5E995eDg1lPn7lUB7TaEKHeiZ4hF
ZIdowLp6vngIAiR/BPAbLANp6FJ0Y7VrDTdOotJwU6vcW2Q0NVdFOsQp1GdaFU48
+x+h9VP442jDk36wz/xPQw7DvrJ/zfcr+x0mn2vkVShrf4V41RnIh7ATj+SGr764
ilpT4np12S1MLDfPZaAM5OcIoaQQdH6l8C6O8FNoS1WEE0yk3s6RG54Yodjk2cxb
9NaOBAbqiKwcg+TRPoBLDcFQrI8gEeq+8m6GX3Q5bmC6IOZh+LzvTkqYuoQKWjJy
+cTFZwlEYrqe5ab+XLGEVnZSr6Xdplt4QGYVAgMBAAECggEAFHlGPadjBUuWuHoJ
VaqrDkVEY+eLbp9cYxenIoFJAH/3zrtewHZeVJ69Qb9HdF+viPY/m5PMzZ8LCXoA
4MciiGQl7/Qm7efDcjvelCgbamuuUV5gx0vfsIGwT5hn8U6fnyfNEsbOThAFWPPZ
c5kN08zN3c4TDH3JVCun4HIup6xjcpVZmLZ810tPNF+dIVhDFq/k+GZi84mWdDlC
2gatmyyGcQk/K7syUJGidGrkUx5uiVPqB8jQcbBD+yGWRWOcJBHYiD6/iR/Xf1tU
j4TGbGrw+3hldQpj0y4A333CjadfthEMPM3GEGauui+esSOuNTyv3pAXB/vD6gaw
kUP1QQKBgQCvA6U7pwbr2g+oZveT2ZPF9TNuqeL0ZGrhxvZiW8aZArl1GJf9F08P
OPGpQtKPyb/LRLx+/fYPBdRyo/fxT8vn/qlozcgC2VkM38mvo1JT9sLA1zC7+dGW
6NG7vd5Kinx1zf57ZntsshVDtU5OsWz0e84+gWbJMNQl62GhT6eYxQKBgQCEkaN4
FlYVcikqWpGv1z4FLSknyO8/6xUO215Gj/V5XX2UbroRdgu0aHx5ZKv8GZy83Hic
uQ4CXbkalrxrEbay6DtE/7beViMrHbPEk91ILCmi7rXEwAkpTRjBblO6WLw5nMeS
6YevB/sJGgK3ndISIhbXqSR083fMjRXxizxNEQKBgB1UaBF6KV/HOI6JSv5dHTW0
pPzrvZwlZAAYXtwW805bNYdZKseAa4Bqk51BFJaCQgEnsMs7dkwINxVLQ1cOf2aO
uAdXTSilEhZlPqCPp1Fo7jRxG7q04BoBRHzJEgK9+KOpdPdrMk4B/ixNqyTm+RJz
2RmCiB45b6Q3MTJebgDdAoGAKtXENDFDb3qLwpuxJ1gtOg71lV1fsQ+MplhLhhFO
CAElaqi8hpfWAF0imzbkO3krI/s5HarN1AXkoarDBvWaSwSu4gSeWgkLJAtc4s/E
WGWsrnDjWseSLj4zGY1EkJnLd/Ioqle699tkSfSVnC7qieFp8BU/Kmrg8r9Dr+CF
g4ECgYEAhReKN82QHB3zARp58KbX83DBRD2uvG4KYeL6bnfGVwPqqwU1Bd8wp7g/
NtfNsQyfOsssfAigxiK0Sck9XzZzQFpMZaeWn9rTB1YYTRnv3qqig6wDgM3aQRCn
/ybLpMhiIv9LQ4ZWD14J3OJjl/RVVPWJQbhC/UDIxkklBNcgfvs=
-----END RSA PRIVATE KEY-----
`

	PEM_ALI_PUB_KEY = `
-----BEGIN PUBLIC KEY-----
MIIBITANBgkqhkiG9w0BAQEFAAOCAQ4AMIIBCQKCAQB9UUWaQISbxJ8+K7BMjk+m
M0ZQ1kR8ejNAem2C6Qh1aihAyd6ZfiupqtCViifUgYBuPewO+qbuvtTzA/MCdgLh
7IIf9KBV8Z4Dko6egMxA3fsoeQDsT2B++jI8/6IgZqTWus3OnR/jEtyLwPUYn+A4
O7nrya9sNvjjkfbGZyMXlGLC+46wcnoanXikMNndVrtunE+Bbf/01EargHU95kB+
r33mFapBEkRZOWVz42qkvYvvssLdbkstUi3g1/BuIca8t+HcrqC64x3NLFJKG3g2
xX89Vfv4gZVjSY3JvQGTvkf2PZFM6uBqYNcbGBq2IaxKFsLIvyZcIeFOP4iAM8NX
AgMBAAE=
-----END PUBLIC KEY-----
`
)

func TestAlipay(t *testing.T) {
	_, err := alipay.NewClient(
		APP_ID,
		SELLER_ID,
		PEM_MY_PUB_KEY,
		PEM_MY_PVT_KEY,
		PEM_ALI_PUB_KEY,
	)
	if err != nil {
		panic(err)
	}
}
