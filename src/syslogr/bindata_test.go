// Code generated by go-bindata.
// sources:
// bindata/ca.crl
// bindata/ca.crt
// bindata/ca.key
// bindata/syslogr.crt
// bindata/syslogr.csr
// bindata/syslogr.key
// DO NOT EDIT!

package main_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)
type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _caCrl = []byte(`-----BEGIN X509 CRL-----
MIICfDBmAgEBMA0GCSqGSIb3DQEBCwUAMA0xCzAJBgNVBAMTAmNhFw0xNzA0MTEy
MTE5MzhaFw0yNzA0MTEyMTE5MzhaMACgIzAhMB8GA1UdIwQYMBaAFEOU6FuUWJw+
1GVsx8xeHgqRwLGvMA0GCSqGSIb3DQEBCwUAA4ICAQA+AVGnxfvg44H/ht3bQuu7
cfPTavxs+/T33/wCgxgBJnRSmfcCGCiekDCt2Cgo2Ycjm1w2eqLGyVICn4xSnnao
6MT/8njuVius9gBgqYqVC5jPI1L15FyXcqlT2etRj3RGDdKYE46GZXSSNORDgQFd
1HLssGFAnh3riBQj3TshtquT1RXLHZXcihv1ogPDSXIn4FOOrouM+hEfoI+77xBE
3RAHS3S6kJ4pwvJJeVUy4gSBy+zUPNeOpuvA/8fbTnldUpxSRk6InaHrBQTEC2Ue
mey6jPmHAIv85WHEORZPONGN0WfXbdgjWRQVHSHJ1fjZO2KSL4FMrKHynQaofVPQ
twfcf2du9jLzgIF2Rcwfky8hC4wbfCiyridmQDyCTZr7w8uiL4BmKY3xVN+bXSBy
afTIK0DVdsHZiKvrBZgCjtOTBeoJaswsvf9AZ9xkkdv3hECtnz7IKnechGyRocZf
ZYaPFs9U0yw2kk68AkyEf+ZY+UdGx3CbVMu3hkOIzAk0eRTPSPSvchAdoIrad0AT
jO/ueYUbdVFhEvlvKggbla0SizgXZXuoc2+wbiqFYl4sPZEn3m63vWnKZnBCdRQ2
EWkKAiLJBcKe+9b145INyzBU+t806S1PYDup75c1DIJIqji/SD28GVAGaYkT3QTk
T3XXH2wsqJE5arphyJ/cLA==
-----END X509 CRL-----
`)

func caCrlBytes() ([]byte, error) {
	return _caCrl, nil
}

func caCrl() (*asset, error) {
	bytes, err := caCrlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ca.crl", size: 918, mode: os.FileMode(292), modTime: time.Unix(1491945578, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _caCrt = []byte(`-----BEGIN CERTIFICATE-----
MIIE2jCCAsKgAwIBAgIBATANBgkqhkiG9w0BAQsFADANMQswCQYDVQQDEwJjYTAe
Fw0xNzA0MTEyMTE5MzVaFw0yNzA0MTEyMTE5MzhaMA0xCzAJBgNVBAMTAmNhMIIC
IjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEArHN/mNV9yn5Nr9PayF/qfz35
vlFcZHBMKSYI+SXJ0kfoYAldN9fer7lS3H0C1q0d6526vpuWoFz5+DlQtn91CxhJ
jx5yzjBQc0Uuzw4uVHDZSiPkye1vngZzcPnmPJx1GAgCUFCobURcljWqRIIzM5cc
v5pMlxOaUNkxFf65yW3vtCld/p+L3kCp6/duuPU6HiwRoDZ1oV7lwE+BR5smCa/a
r5VF6VSv52pITtjpZAla4feUCqCPR3a3hg8YBxO2DklREL7JyDtpaQseBi4J7SFl
THDpoylfB/6yK/lPerfKqqDwfJLUpRFCoUSdntkvh2n7FXrfwOxVVDNJ6S+tT8Sx
SNWZKbpie82d0xjFb9D3G81bdEE3CuFKHVDvs5Vpj+QPHjXqf0t4koQmbk4ee9NY
K3Dy4bKdanORrQtxUQx6u/sLxIHzGXbkDjV0EpIjozwVV7zEK8rB2+NI4nEaa9ca
g2HX/kAtFV54N1ZE7fChN7TXXbCxL/shr5shkfjU9xuUma+6UJRBv+x+mE7iXZWU
3ZK/WY1/gPKSAqBYD35OtLr26chU67iJvxTNwaimXKpTAFJdnt5Jjs+oBuR58l4p
lk6U60HuAkKln+IHA8vR38sHnlAMs9NXMoiuhIeFzcLeytdgARUBbfvGR0lUVIbH
i0pZSaJHov1FJizZ0bcCAwEAAaNFMEMwDgYDVR0PAQH/BAQDAgEGMBIGA1UdEwEB
/wQIMAYBAf8CAQAwHQYDVR0OBBYEFEOU6FuUWJw+1GVsx8xeHgqRwLGvMA0GCSqG
SIb3DQEBCwUAA4ICAQBtcgP1BjJhwCC947jtgc8rQC7DRQCuhZ1ouYdvnU1uQMYq
GE0/Jxs0VeSwrsDtpe3XoP3KMnlE/66tbUJ+xFYq1N0wZITZ/s0U3HBGwAvchd/P
ZzDhbIOfEU4Xvu0o5DAMrBbmdGdQ9SgPJpeqL4CalQy/HjzXGEePErQO+2LkNiEz
GRMtJWXIh70jGYXng2P/+lgBos8M+TaDOUq4lgqx8Z2Ihiwm4Tqn22TA5NkmqheT
eHE7bfXPAMfO7H48NVMxSLemGar3ZIQQLl7vSXhDOkrDane7jLbbaiPNwuV4mdJ2
Aj3OK0N4IPog0HCI1YiEwod0biO7mgHaeacBvrXCZGaY1JhSuc70/AUdqhUlg9MT
oWtW9bYKMrF6w35IHCFg+F0QxFf3CH5rbqSbRQyvBAaAsWh5P19VqUoISMQrGvco
kVlQhuWZwkJQLvvQZeguJ64isUYn2YTjp+qyeTwRLAyAy/wCFr0xE8Dl277Ml+fS
d72emjrPWCEC1fl5mpneWPNQcSYlDUihF152pj1HlEWUbw2A5Ezsa4hSb1O0gBtC
izdbCLrRKBNzj/Y7VFjwNUsiAd8pCWdiQlpXA+5Zxi/O6EuKUh65ATAcCH+DHsQq
7yo79okyaMJPMe71penZ3ECGiDxOiJtOCI9uAjoXcK/utVRCa9gbPMpe58JiTg==
-----END CERTIFICATE-----
`)

func caCrtBytes() ([]byte, error) {
	return _caCrt, nil
}

func caCrt() (*asset, error) {
	bytes, err := caCrtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ca.crt", size: 1744, mode: os.FileMode(292), modTime: time.Unix(1491945578, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _caKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----
`)

func caKeyBytes() ([]byte, error) {
	return _caKey, nil
}

func caKey() (*asset, error) {
	bytes, err := caKeyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ca.key", size: 3243, mode: os.FileMode(288), modTime: time.Unix(1491945578, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _syslogrCrt = []byte(`-----BEGIN CERTIFICATE-----
MIIEGjCCAgKgAwIBAgIQc/kjpG6IxOyhPIiWD5DTzjANBgkqhkiG9w0BAQsFADAN
MQswCQYDVQQDEwJjYTAeFw0xNzA0MTEyMTE5NDZaFw0xOTA0MTEyMTE5NDZaMBIx
EDAOBgNVBAMTB3N5c2xvZ3IwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIB
AQCrS5kkG92dK40qxA6BTARnRdvMxfh6LMKbu0h9bO8kn8KceGzCZBul64c6k+Ce
rgCcN7aYo3Dk3qmt4efTD/vNMhZse+VDQk8PeOWhyyojrlzkY+lxMkRhL7hfYdwF
1rKZwu4iGMzKNGL7iQgGXyMe6rJ0BHWD5/flq1YxDVQlcANPbu1nkLuw6v7U2XZJ
rgt4czNeB847Y/4/j3x93i0NdCMhWDgRNQzDVDmKStbg/kBIDebxcVgnYhYQJso/
qm3IEY3qQAN6udf/XR0uK0QOBhjrq+Gg/CZ0LNJeFCwcThWdcKANbDs9/Tp49Dhs
DhTylGMrcuX8NmNWNutHDwhVAgMBAAGjcTBvMA4GA1UdDwEB/wQEAwIDuDAdBgNV
HSUEFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIwHQYDVR0OBBYEFCXdPjsEv2Oq7MlJ
8tZOvfsOCdPAMB8GA1UdIwQYMBaAFEOU6FuUWJw+1GVsx8xeHgqRwLGvMA0GCSqG
SIb3DQEBCwUAA4ICAQCsNJ7FmZ4COuFUwZt1GVG7bniccltQNjpUi0krQufBf3u7
kchvacTKoviVGLjskIFYTAsNhgODymDh5IJrmX/cLTrCk4Fcu2ntrM91KVYwCAPz
cKl4OZcKWh+6+9NfNJrZJ3ChuarEqZNbK8wlK48TkRB36zNIyLLLbeDZiPZwhGmM
Ip33KEvi08xk1j1FRIWvbejK/2EC88z0VpAHRMNSJt2nWj1/HVC8wA+drW7zVBSx
jUUOUvRpBxb4A7prdmJlvWaKq+f2Nsaet4BB0PU53encaIoBtbbJYQmd5HjEbSED
OGDDz1NVJLHEsAbfMhe9Y3inKk8+jylkydbGycVc/FwDZg8kyXpEMBVVDjUl1kO6
Vj5KM1Qh4JbXuCbvypCXFnpdrcoxqweNPIMTZeR9AQiaNd0hi+7oXPrCu4SKk+Gb
3/dOobvqJCY0IYW3eO5yURTDvJkFhtOBVrSZLWH5QVUNsiuIh+70kWgKGpbJNnep
7M1JHrQDw1rvNSDVc+VvFdlNE1qgQM62g+vcBmyGH2AyPodZ+gEAfGUkB4ho+6lY
X3vDNDVKpM53HNARo+J16v+CIMNQ9jgPLz84wEF2yCPoIG010vQLsmQFTlbcMKuF
CcVAbdbpcMqdisWaLeho7dj4vsb8ufSxhq+GgC2ckDKEpPlwk+DF7fcRNrA/Rg==
-----END CERTIFICATE-----
`)

func syslogrCrtBytes() ([]byte, error) {
	return _syslogrCrt, nil
}

func syslogrCrt() (*asset, error) {
	bytes, err := syslogrCrtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "syslogr.crt", size: 1484, mode: os.FileMode(292), modTime: time.Unix(1491945586, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _syslogrCsr = []byte(`-----BEGIN CERTIFICATE REQUEST-----
MIICVzCCAT8CAQAwEjEQMA4GA1UEAxMHc3lzbG9ncjCCASIwDQYJKoZIhvcNAQEB
BQADggEPADCCAQoCggEBAKtLmSQb3Z0rjSrEDoFMBGdF28zF+Hoswpu7SH1s7ySf
wpx4bMJkG6XrhzqT4J6uAJw3tpijcOTeqa3h59MP+80yFmx75UNCTw945aHLKiOu
XORj6XEyRGEvuF9h3AXWspnC7iIYzMo0YvuJCAZfIx7qsnQEdYPn9+WrVjENVCVw
A09u7WeQu7Dq/tTZdkmuC3hzM14Hzjtj/j+PfH3eLQ10IyFYOBE1DMNUOYpK1uD+
QEgN5vFxWCdiFhAmyj+qbcgRjepAA3q51/9dHS4rRA4GGOur4aD8JnQs0l4ULBxO
FZ1woA1sOz39Onj0OGwOFPKUYyty5fw2Y1Y260cPCFUCAwEAAaAAMA0GCSqGSIb3
DQEBCwUAA4IBAQCfz9Kh5V67UbkF0AM55JAi96f2Gqv/JCocoQsxYykLbi6jd/Yx
RKjtLwmY/mG0bY787FuGlYKATzpHSnMlX/uPNMSQ3RYHSpxsEM/R/I0HNBm+3aUo
sGdtxdcYarlspolfP0wU1my4nBC2hRa28R6jStEVztsHzXOf+YfUSvogIoQDxU5i
9P60z4l6P3uzElTHKBRRkT2MZHweoGoSq+XWxhdIGmo3mjZKHiP7ekZ+UKBN5Ncq
xSBgmjTJDxchFXHnyYKvJS5ls2xbGqPpTFoE9vzYW6aFpG1jBvgfyy96vPkSJ7l3
qEAineIlIjUsshKr9PcnO8nJP2zCGTy++75H
-----END CERTIFICATE REQUEST-----
`)

func syslogrCsrBytes() ([]byte, error) {
	return _syslogrCsr, nil
}

func syslogrCsr() (*asset, error) {
	bytes, err := syslogrCsrBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "syslogr.csr", size: 887, mode: os.FileMode(292), modTime: time.Unix(1491945485, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _syslogrKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAq0uZJBvdnSuNKsQOgUwEZ0XbzMX4eizCm7tIfWzvJJ/CnHhs
wmQbpeuHOpPgnq4AnDe2mKNw5N6preHn0w/7zTIWbHvlQ0JPD3jlocsqI65c5GPp
cTJEYS+4X2HcBdaymcLuIhjMyjRi+4kIBl8jHuqydAR1g+f35atWMQ1UJXADT27t
Z5C7sOr+1Nl2Sa4LeHMzXgfOO2P+P498fd4tDXQjIVg4ETUMw1Q5ikrW4P5ASA3m
8XFYJ2IWECbKP6ptyBGN6kADernX/10dLitEDgYY66vhoPwmdCzSXhQsHE4VnXCg
DWw7Pf06ePQ4bA4U8pRjK3Ll/DZjVjbrRw8IVQIDAQABAoIBAFr7H07fghjtveAe
HDounxQfNbyQ1gZGaeL/WWRNDMPeGyayi8nAFDNPYxcI3kBJ2UcgDFsMTHUzktop
Z9Fh5vM5DFH4iG/t80Ibi5Qg7bPf87TXIHGgKlOYXgxDVPLA80HaGFOB6pvnsT86
rEl+FJBiRgB7MdOqzK9vDqulEbGD67q8gf6x9zoILmxL8MmkUX1Y04xnWcKBiS2a
8txZSf1ZDB00vZzPft3zlBKmKEmdIdnhr1D8O7WtKdluFnWfskI3oE8Ps/Vx+Dxv
jc73+XM9EXMngAyjprYA7H5dnIVSzJ1+/ijORlMbmmm2NW/KIwyoxatXxsnp6aUp
cC2iDc0CgYEA0WwxnkMBrNvy3hYYHI2upYtl03aslXWhuMPM0QAU6rvRQSBeoLOV
DE3ywNQEcXn8Xmc2vGvE/FAhjjNvU0PExkZ0kBOvMG7ptYTMg+XerJaMKPqRF3Ks
/2jBglVL5UA6StsfgonhfeJiaAkdwyuezdrmutRkUbuVets+WSxk+wsCgYEA0WSQ
sg7qQNluLX6DAakxDeGfA9UB/yqX2JmnvA0jLdbprE6s7DtO+SFWqQ0KbbqwsYe8
3YByla049YqoEMCXA64MrAIANtv3N+n2ulCLitwS2GYP9mqtLRtwluY846nq9fMK
8wV/q5n4hVaSGt2q/yzUuqx9bmqO7CZQSq9TJh8CgYB175MlyCBqY1crYI/ljJh8
27rVUnCwnpUbgxCV5pDg/DJEleEUaO2YB1Gc7AkwinzoAQIfLCW73bh351lbL7/h
1Q2RkbRH9z4gHhA6ezpiQUnfTfRlmwv3rdvD9RsPmJL7Utk030cgaFv4sKGVJryA
uzTSpAd2y8fLbyp2d3zKrwKBgQDKw10ybyr73RDJ0SY/J30u+fonlc63LqGXx0Mb
4ITeUKw6hH0CzHOu30+xp8UfeJa9crm6rDVJJ1JuwvP4NOaHU7VFozJd7Sc5579Z
r8FQl8dP8ZUngylq8pVWKmFv/AxgWheSORLmMtTrGWelyF1beCgPFBTqJRl9J6S8
jYYw9QKBgD4UeYwWv5FNTiOfwkMhAJds7n0pdE97htowP4UunYcOLjR+2vNJwKx8
4XPe6IiJpXe+Tnvkt+5h/TqpBAXKq1YW7jDC9QmlUzHU6mtkgTY8/dObYemUfwO6
LlsRdiJEkC6gzML5LlIZjXUXq4wr/v2+kk2PshIbpi5grjVeKo+j
-----END RSA PRIVATE KEY-----
`)

func syslogrKeyBytes() ([]byte, error) {
	return _syslogrKey, nil
}

func syslogrKey() (*asset, error) {
	bytes, err := syslogrKeyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "syslogr.key", size: 1675, mode: os.FileMode(288), modTime: time.Unix(1491945485, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"ca.crl": caCrl,
	"ca.crt": caCrt,
	"ca.key": caKey,
	"syslogr.crt": syslogrCrt,
	"syslogr.csr": syslogrCsr,
	"syslogr.key": syslogrKey,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"ca.crl": &bintree{caCrl, map[string]*bintree{}},
	"ca.crt": &bintree{caCrt, map[string]*bintree{}},
	"ca.key": &bintree{caKey, map[string]*bintree{}},
	"syslogr.crt": &bintree{syslogrCrt, map[string]*bintree{}},
	"syslogr.csr": &bintree{syslogrCsr, map[string]*bintree{}},
	"syslogr.key": &bintree{syslogrKey, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

