package bangun_datar

func LuasSegitiga(Alas int64, Tinggi int64) int64 {
	return (Alas * Tinggi / 2)
}

func KelilingSegitiga(Alas int64, Tinggi int64) int64 {
	return Alas * 3
}

func LuasPersegi(Sisi int64) int64 {
	return Sisi * Sisi
}

func KelilingPersegi(Sisi int64) int64 {
	return Sisi * 4
}

func LuasPersegiPanjang(Panjang int64, Lebar int64) int64 {
	return Panjang * Lebar
}

func KelilingPersegiPanjang(Panjang int64, Lebar int64) int64 {
	return (Panjang * 2) + (Lebar * 2)
}

func LuasLingkaran(JariJari int64) int64 {
	return JariJari * JariJari * 22 / 7
}

func KelilingLingkaran(JariJari int64) int64 {
	return JariJari * 11 / 7
}
