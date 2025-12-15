module git.techpartners.asia/gateway-services/payment-service

go 1.25.4

require (
	github.com/bytedance/sonic v1.14.2
	github.com/getsentry/sentry-go v0.40.0
	github.com/getsentry/sentry-go/fiber v0.40.0
	github.com/gofiber/fiber/v2 v2.52.10
	github.com/redis/go-redis/v9 v9.17.2
	github.com/spf13/viper v1.21.0
	github.com/techpartners-asia/balc-api-go v1.0.0
	github.com/techpartners-asia/golomt-api-go v0.0.15
	github.com/techpartners-asia/monpay-go v1.0.0
	github.com/techpartners-asia/pocket-go v0.0.0-20250109090209-90886b51b423
	github.com/techpartners-asia/qpay-go v1.0.2
	github.com/techpartners-asia/simple-go v1.0.2
	github.com/techpartners-asia/storepay-go v1.0.0
	github.com/techpartners-asia/tokipay-go v1.0.0
	google.golang.org/grpc v1.77.0
	google.golang.org/protobuf v1.36.10
	gorm.io/driver/postgres v1.6.0
	gorm.io/gorm v1.31.1
)

require (
	github.com/andybalholm/brotli v1.1.1 // indirect
	github.com/bytedance/gopkg v0.1.3 // indirect
	github.com/bytedance/sonic/loader v0.4.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cloudwego/base64x v0.1.6 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/fsnotify/fsnotify v1.9.0 // indirect
	github.com/go-viper/mapstructure/v2 v2.4.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/schema v1.4.1 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.6.0 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/klauspost/cpuid/v2 v2.2.9 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/pelletier/go-toml/v2 v2.2.4 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/sagikazarmark/locafero v0.11.0 // indirect
	github.com/sourcegraph/conc v0.3.1-0.20240121214520-5f936abd7ae8 // indirect
	github.com/spf13/afero v1.15.0 // indirect
	github.com/spf13/cast v1.10.0 // indirect
	github.com/spf13/pflag v1.0.10 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.57.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	go.yaml.in/yaml/v3 v3.0.4 // indirect
	golang.org/x/arch v0.0.0-20210923205945-b76863e36670 // indirect
	golang.org/x/crypto v0.46.0 // indirect
	golang.org/x/net v0.47.0 // indirect
	golang.org/x/sync v0.19.0 // indirect
	golang.org/x/sys v0.39.0 // indirect
	golang.org/x/text v0.32.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251022142026-3a174f9686a8 // indirect
	resty.dev/v3 v3.0.0-beta.3 // indirect
)

replace git.techpartners.asia/gateway-services/payment-service/pkg/proto/payment => ./pkg/proto/payment
