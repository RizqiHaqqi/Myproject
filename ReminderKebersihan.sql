CREATE TABLE

CREATE TABLE Admin (
    admin_id INT PRIMARY KEY,
    name VARCHAR(50),
    email VARCHAR(100),
    password VARCHAR(50)
);

CREATE TABLE Users(
    user_id INT PRIMARY KEY,
    name VARCHAR(50),
    email VARCHAR(100),
    password VARCHAR(50),
    birth_date DATE,
    profile_picture VARCHAR(255),
    earned_points INT
);

CREATE TABLE Tasks (
    task_id INT PRIMARY KEY,
    title VARCHAR(100),
    task_description VARCHAR(1000),
    task_poin INT,
    start_date DATE,
    end_date DATE
);


CREATE TABLE Submission (
    Submission_id INT PRIMARY KEY,
    User_id INT,
    Task_id INT,
    Video VARCHAR(255),
    Picture VARCHAR(255),
    FOREIGN KEY (User_id) REFERENCES Users(User_id),
    FOREIGN KEY (Task_id) REFERENCES Tasks(Task_id)
);

CREATE TABLE Reward(
    reward_id INT PRIMARY KEY,
    reward_title VARCHAR(100),
    description VARCHAR(1000),
    needed_poin INT,
    reward_type VARCHAR(100)
);

1. Views
Views adalah objek database yang menyajikan data dari satu atau beberapa tabel dalam bentuk yang lebih mudah diakses dan dianalisis. Views juga dapat digunakan untuk menyederhanakan query yang kompleks.

View untuk Laporan Pengguna dan Poin

CREATE VIEW UserPoints AS
SELECT 
    U.User_id,
    U.Name,
    U.Email,
    U.BirthDate,
    U.Profile_Picture,
    SUM(S.Task_Point) AS Total_Points
FROM User U
JOIN Submission S ON U.User_id = S.User_id
JOIN Task T ON S.Task_id = T.Task_id
WHERE S.Status = 'Valid'
GROUP BY U.User_id, U.Name, U.Email, U.BirthDate, U.Profile_Picture;

2. Functions
Functions adalah blok kode yang dapat dipanggil dengan nama tertentu dan mengembalikan nilai tertentu. Functions sering digunakan untuk perhitungan atau manipulasi data tertentu.

Function untuk Menghitung Poin Total Pengguna

CREATE FUNCTION GetUserTotalPoints(userId INT)
RETURNS INT
BEGIN
    DECLARE totalPoints INT;
    SELECT SUM(T.Task_Point) INTO totalPoints
    FROM Submission S
    JOIN Task T ON S.Task_id = T.Task_id
    WHERE S.User_id = userId AND S.Status = 'Valid';
    RETURN totalPoints;
END;
