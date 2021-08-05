# Naming Conventions trong Go

## Method

-  Nếu muốn khai báo một Method sử dụng được cả bên ngoài packet thì tên bắt đầu bằng chữ hoa
- Nếu chỉ muốn một Method sử dụng bên trong packet thì nên đặt theo kiểu camelCase
```sh
package awesome
type Awesomeness struct {}
// Do is an exported method and is accessible in other packages
func (a Awesomeness) Do() string {
 return a.doMagic("Awesome")
}
// doMagic is where magic happens and only visible inside awesome 
func (a Awesomeness) doMagic(input string) string {
 return input
}
```

## Đặt tên interface

Tên interface nên bắt đầu bằng phương thức của interface đó + hậu tố er
```sh
MethodName + er = InterfaceName
EX:Formatter, CloseNotifier 
```


## Getter setter
Go khuyên khích không nên dùng

.
# Những quy tắc cần nhớ
Code của chương trình phải đọc hiểu được 
### Đăt tên quá ngắn
Không miêu tả dược hành động của phương thức
### Tên viết tắt
Miễn là hiểu được,tuy nhiên phạm vi sử dụng càng rộng thì cảng phải chi tiết

### Các từ đặc biệt không nên thay đổi kiểu ký tự
```
ID->userID
DB->loadDB
```
### Độ dài câu lệnh
Không nên quá dài

# Tổng kết
Các convension là cần thiết để quá tình viết và đọc dễ dàng và logic,tuy nhiên cũng ko cần quá cứng nhắc.

