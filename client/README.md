1. Terdiri dari tiga komponen utama, yaitu frontend,backend, dan database, front end digunakan untuk menampilkan daftar task, mengisi form, serta melakukan aksi seperti menambah, mengubah, dan menghapus task. Frontend berkomunikasi dengan backend melalui HTTP request menggunakan fetch API. Backend dibuat dengan Gin yang bertugas menerima request dari frontend, terdapat  autentikasi sederhana menggunakan middleware, lalu menghubungkan menggunakan GORM untuk berinteraksi dengan database MySQL.

2. skema menggunakan struct dari seperti berikut

type TaskRequest struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	Status      string    `json:"status" binding:"required,oneof=todo inprogress done"`
	DueDate     time.Time `json:"due_date" binding:"required"`
} title dan status diatur sebagai required sementara description tidak dan due date si set sebagai time agar sistem dapat menyimpan tanggal dengan baik.

3. endpoint diatur pada file route dengan endpoint yang sama yaitu api/task , hanya dibedakan metho nya sesuai dengan fungsi yang menangani nya di controller antara lain
get - tanpa id : digunakan untuk fetch semua data task yg ada
get/:id : digunakan untuk fetch task tertentu berdasarkan id
post : digunakan untuk membuat task baru dan menyimpannya di database
patch: digunakan untuk memperbarui data status task sesuai id
delete : digunakan untuk menghapus data task sesuai id

4. menurut saya semua bagian cocok baik di go dan node.js namun khusus dibagian migration saya lebih suka go karena bisa automigrate , simple
5. strategi testing mungkin dibagian api, diuji dibagian middleware

