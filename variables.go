package main

// Целочисленные типы
var a int8 = -1   // целое число от -128 до 127 и занимает 1 байт
var b uint = 2    // целое число от 0 до 255 и занимает 1 байт
var c byte = 3    // byte - синоним типа uint8
var d int16 = -4  // целое число от -32'768 до 32'767 и занимает 2 байта
var f uint16 = 5  // целое число от 0 до 65'535 и занимает 2 байта
var g int32 = -6  // целое число от -2'147'483'648 до 2'147'483'647 и занимает 4 байта
var h rune = -7   // целое число от -2'147'483'648 до 2'147'483'647 и занимает 4 байта
var j uint32 = 8  // целое число от 0 до 4294967295 и занимает 4 байта
var k int64 = -9  // целое число от –9 223 372 036 854 775 808 до 9 223 372 036 854 775 807 и занимает 8 байт
var l uint64 = 10 // целое число от 0 до 18 446 744 073 709 551 615 и занимает 8 байт
var m int = 102   // целое число со знаком, которое в зависимости о платформы может занимать либо 4 байта, либо 8 байт. То есть соответствовать либо int32, либо int64
var n uint = 105  // целое беззнаковое число только без знака, которое, аналогично типу int, в зависимости о платформы может занимать либо 4 байта, либо 8 байт. То есть соответствовать либо uint32, либо uint64

//Числа с плавающей точкой
var fa float32 = 18 // число с плавающей точкой от 1.4*10-45 до 3.4*1038(для положительных). Занимает в памяти 4 байта
var ga float32 = 4.5
var da float64 = 0.23 // число с плавающей точкой от 4.9*10-324 до 1.8*10308 (для положительных) и занимает 8 байт
var ea float64 = 2.7

// В качестве разделителя между целой и дробной частью применяется точка

// Комплексные числа
var fb complex64 = 1 + 2i  // комплексное число, где вещественная и мнимая части представляют числа float32
var gb complex128 = 4 + 3i // комплексное число, где вещественная и мнимая части представляют числа float64

// Тип bool
var isAlive bool = true
var isEnabled bool = false

// Строки
var name string = "Том Сойер"

// Кроме обычных символов строка может содержать специальные последовательности(управляющие последовательности), которые начинаются с обратного слеша \. Наиболее распространенные последовательности:
// \n - переход на новую строку
// \r - возврат каретки
// \t - табуляция
// \" - двойная кавычка внутри строк
// \\ - обратный слеш

// Значение по умолчанию
// Если переменной не присвоено значение, то она имеет значение по умолчанию, которое определено для ее типа. Для числовых типов это число 0, для логического типа - false, для строк - "" (пустая строка)

// Неявная типизация
// При определеннии переменной мы можем опускать тип в том случае, если мы явно инициализируем переменную каким-нибудь значением:
var nameA string = "Том"

// В этом случае компилятор на основании значения неявно выводит тип переменной. Если присваивается строка, то соответственно переменная будет представлять тип string, если присваивается целое число, то переменная представляет тип int и т.д.
var nameB = "Том"

// При этом стоит учитывать, что если мы не указываем у переменной тип, то ей обязательно надо присвоить некоторое начальное значение. Объявление переменной одновременно без указания типа данных и начального значения будет ошибкой.
//var nameC // !Ошибка

// Надо либо указать тип данных (в этом случае переменная будет иметь значение по умолчанию)
var nameD string

// Либо указать начальное значение, на основании которого выводится тип данных
var nameE = "Tom"

//Либо и то, и другое одновременно:
var nameF string = "Tom"

// Неявная типизация нескольких переменных
var (
	nameG = "Tom"
	age   = 27
)

// Или так
var nameH, ageA = "Tom", 27
