-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jan 20, 2026 at 06:44 AM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `pemilihan`
--

-- --------------------------------------------------------

--
-- Table structure for table `kandidat`
--

CREATE TABLE `kandidat` (
  `id_kandidat` int(11) NOT NULL,
  `nomor_urut` int(11) NOT NULL,
  `nama_ketua` varchar(100) NOT NULL,
  `nama_wakil` varchar(100) NOT NULL,
  `visi` text NOT NULL,
  `misi` text NOT NULL,
  `foto` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `kandidat`
--

INSERT INTO `kandidat` (`id_kandidat`, `nomor_urut`, `nama_ketua`, `nama_wakil`, `visi`, `misi`, `foto`, `created_at`) VALUES
(2, 2, 'Rizky', 'Fajar', 'OSIS yang inovatif dan disiplin', '1. Program digital\n2. Ketertiban sekolah', NULL, '2026-01-19 00:59:18'),
(3, 1, 'tampling', 'sisi', 'pribadi yang baik', 'ikut sholat magrib', NULL, '2026-01-19 02:54:04');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id_user` int(11) NOT NULL,
  `username` varchar(50) NOT NULL,
  `email` varchar(100) NOT NULL,
  `password` varchar(255) NOT NULL,
  `role` enum('user','admin','','') NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `status` enum('aktif','sudah_voting','nonaktif') DEFAULT 'aktif'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id_user`, `username`, `email`, `password`, `role`, `created_at`, `status`) VALUES
(19, 'ayam', 'ayam@gmail.com', '$2a$10$UHbgHJAOVY3UyjT7q3x3nuR88kwdN7qutrKB/ZXW/R4b1zkO0RQhW', 'admin', '2026-01-13 03:17:15', 'aktif'),
(20, 'bebek', 'bebek@gmail.com', '$2a$10$aa6ZE4VqyFitCGcPrEyEQeWQImMOOoXy618CMF7b37Zd8I2Ee/yqu', 'user', '2026-01-13 03:37:26', 'aktif'),
(21, 'yam', 'yam@gmail.com', '$2a$10$lNJbUHE8fJmMQISIxK63puS0aoL8GIReNjV4PI7q31ecegtcIQU72', 'admin', '2026-01-14 00:29:31', 'aktif'),
(22, 'gagak', 'gagak@gmail.com', '$2a$10$2G6hcK.MQ1l9VjUBXHW9X.5LqJ9XGUAlSL91rD8UOEA3G0usTs7s6', 'user', '2026-01-20 03:28:34', 'aktif'),
(23, 'guss', 'guss@gmail.com', '$2a$10$OVEWFQdP5MRWtnQChGaK0OWaAyf8DhuXZXE29smTt4FlEhdLYeFdW', 'user', '2026-01-20 03:58:53', 'sudah_voting'),
(24, 'king', 'king@gmail.com', '$2a$10$JJHv7fo1MjSyFRUXuR8Dj.d4VAGJznyNuKEZs.JzDHuBaz5SUanNi', 'user', '2026-01-20 04:04:46', 'aktif');

-- --------------------------------------------------------

--
-- Table structure for table `voting`
--

CREATE TABLE `voting` (
  `id_voting` int(11) NOT NULL,
  `id_user` int(11) NOT NULL,
  `id_kandidat` int(11) NOT NULL,
  `waktu_voting` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `voting`
--

INSERT INTO `voting` (`id_voting`, `id_user`, `id_kandidat`, `waktu_voting`) VALUES
(1, 23, 2, '2026-01-20 04:26:51');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `kandidat`
--
ALTER TABLE `kandidat`
  ADD PRIMARY KEY (`id_kandidat`),
  ADD UNIQUE KEY `nomor_urut` (`nomor_urut`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id_user`),
  ADD UNIQUE KEY `nis` (`username`),
  ADD UNIQUE KEY `email` (`email`);

--
-- Indexes for table `voting`
--
ALTER TABLE `voting`
  ADD PRIMARY KEY (`id_voting`),
  ADD UNIQUE KEY `id_user` (`id_user`),
  ADD KEY `fk_voting_kandidat` (`id_kandidat`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `kandidat`
--
ALTER TABLE `kandidat`
  MODIFY `id_kandidat` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id_user` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=25;

--
-- AUTO_INCREMENT for table `voting`
--
ALTER TABLE `voting`
  MODIFY `id_voting` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `voting`
--
ALTER TABLE `voting`
  ADD CONSTRAINT `fk_voting_kandidat` FOREIGN KEY (`id_kandidat`) REFERENCES `kandidat` (`id_kandidat`) ON DELETE CASCADE,
  ADD CONSTRAINT `fk_voting_user` FOREIGN KEY (`id_user`) REFERENCES `users` (`id_user`) ON DELETE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
