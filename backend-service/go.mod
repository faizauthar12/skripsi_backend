module github.com/faizauthar12/skripsi/backend-service

go 1.21.0

require (
	github.com/ethereum/go-ethereum v1.12.2
	github.com/faizauthar12/skripsi/cart-gomod v0.0.0-00010101000000-000000000000
	github.com/faizauthar12/skripsi/cart-item-gomod v0.0.0-00010101000000-000000000000
	github.com/faizauthar12/skripsi/customer-gomod v0.0.0-00010101000000-000000000000
	github.com/faizauthar12/skripsi/order-gomod v0.0.0-00010101000000-000000000000
	github.com/faizauthar12/skripsi/product-gomod v0.0.0-00010101000000-000000000000
	github.com/faizauthar12/skripsi/user-gomod v0.0.0-00010101000000-000000000000
	github.com/gin-contrib/cors v1.4.0
	github.com/gin-gonic/gin v1.9.1
	github.com/go-playground/form/v4 v4.2.1
	github.com/joho/godotenv v1.5.1
	go.mongodb.org/mongo-driver v1.12.1
)

require (
	github.com/btcsuite/btcd/btcec/v2 v2.3.2 // indirect
	github.com/bytedance/sonic v1.10.0 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20230717121745-296ad89f973d // indirect
	github.com/chenzhuoyu/iasm v0.9.0 // indirect
	github.com/dchest/uniuri v1.2.0 // indirect
	github.com/deckarep/golang-set/v2 v2.3.1 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.2.0 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.2 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.15.1 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/golang-jwt/jwt/v5 v5.0.0 // indirect
	github.com/golang/snappy v0.0.5-0.20220116011046-fa5810519dcb // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/holiman/uint256 v1.2.3 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.16.7 // indirect
	github.com/klauspost/cpuid/v2 v2.2.5 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/montanaflynn/stats v0.7.1 // indirect
	github.com/pelletier/go-toml/v2 v2.0.9 // indirect
	github.com/shirou/gopsutil v3.21.11+incompatible // indirect
	github.com/tklauser/go-sysconf v0.3.12 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.11 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20201027041543-1326539a0a0a // indirect
	github.com/yusufpapurcu/wmi v1.2.3 // indirect
	golang.org/x/arch v0.4.0 // indirect
	golang.org/x/crypto v0.12.0 // indirect
	golang.org/x/exp v0.0.0-20230811145659-89c5cff77bcb // indirect
	golang.org/x/net v0.14.0 // indirect
	golang.org/x/sync v0.3.0 // indirect
	golang.org/x/sys v0.11.0 // indirect
	golang.org/x/text v0.12.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/faizauthar12/skripsi/cart-gomod => ../cart-gomod
	github.com/faizauthar12/skripsi/cart-item-gomod => ../cart-item-gomod
	github.com/faizauthar12/skripsi/customer-gomod => ../customer-gomod
	github.com/faizauthar12/skripsi/order-gomod => ../order-gomod
	github.com/faizauthar12/skripsi/product-gomod => ../product-gomod
	github.com/faizauthar12/skripsi/user-gomod => ../user-gomod
)

replace github.com/faizauthar12/skripsi/email-gomod => ../email-gomod
