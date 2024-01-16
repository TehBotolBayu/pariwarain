 CREATE TABLE LaporanMingguan (
     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
     PenyewaanId UUID REFERENCES Penyewaan(id) ON DELETE CASCADE,
     Foto VARCHAR(255),
     createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     deskripsi TEXT
 );

 CREATE TABLE Penyewaan (
     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
     postId UUID REFERENCES Properti(id) ON DELETE CASCADE,
     penyewaId UUID REFERENCES Pengguna(id) ON DELETE CASCADE,
     pemilikId UUID REFERENCES Pengguna(id) ON DELETE CASCADE,
     startRent TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     endRent TIMESTAMP,
     price BIGINT,
     status VARCHAR(255),
     kodeUnik VARCHAR(255)
 );

CREATE TABLE Properti (
    Id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    foto VARCHAR(255),
    judul VARCHAR(255),
    alamat VARCHAR(255),
    harga BIGINT,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP,
    PenggunaId UUID REFERENCES Pengguna(Id) ON DELETE CASCADE
);

CREATE TABLE Pengguna (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    nama VARCHAR(255),
    email VARCHAR(255),
    pw VARCHAR(255)
);