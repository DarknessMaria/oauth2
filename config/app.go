package config

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
type App struct {
    Session struct {
        Name      string `yaml:"name"`
        SecretKey string `yaml:"secret_key"`
        MaxAge    int    `yaml:"max_age"`
    } `yaml:"session"`
    Db struct {
        Default Db
    }
    Redis struct {
        Default Redis
    }
    OAuth2 struct {
        AccessTokenExp int      `yaml:"access_token_exp"`
        JWTSignedKey   string   `yaml:"jwt_signed_key"`
        Client         []Client `yaml:"client"`
    } `yaml:"oauth2"`
}

type Db struct {
    Type     string `yaml:"type"`
    Host     string `yaml:"host"`
    Port     int    `yaml:"port"`
    User     string `yaml:"user"`
    Password string `yaml:"password"`
    DbName   string `yaml:"dbname"`
}

type Redis struct {
    Addr     string `yaml:"addr"`
    Password string `yaml:"password"`
    Db       int    `yaml:"db"`
}

type Client struct {
    ID     string  `yaml:"id"`
    Secret string  `yaml:"secret"`
    Name   string  `yaml:"name"`
    Domain string  `yaml:"domain"`
    Scope  []Scope `yaml:"scope"`
}

type Scope struct {
    ID    string `yaml:"id"`
    Title string `yaml:"title"`
}
