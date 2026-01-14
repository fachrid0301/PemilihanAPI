-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jan 13, 2026 at 04:43 AM
-- Server version: 10.4.24-MariaDB
-- PHP Version: 8.1.6

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
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `username` varchar(50) NOT NULL,
  `email` varchar(100) NOT NULL,
  `password` varchar(255) NOT NULL,
  `role` enum('user','admin','','') NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `username`, `email`, `password`, `role`, `created_at`) VALUES
(1, '12345', '1234@gmail.com', '12345', 'user', '2026-01-07 03:31:52'),
(2, '123', '123@gmail.com', '123@gmail.com', 'user', '2026-01-07 03:31:52'),
(4, 'fachri', 'fachr@gmail.com', '$2a$10$FEY6xU6fj6kTRzDHmEBlhekz5IunSgJ/zgGgmpwZpd4mnkx.bWPgu', 'user', '2026-01-07 07:05:15'),
(6, 'exejar', 'adversaries.of.chaos@gmail.com', '$2a$10$FYn0wIG7AXGnxupJpksDX.dCT41r8p0aA/xMwJH6BES.FENjqKihy', 'admin', '2026-01-13 01:13:49'),
(7, 'haidar', 'haidarfarin@gmail.com', '$2a$10$8duULBWVXg7sJU7tssg1L.NzTpAeu0ngoVzVAg/vi/RsbmvIeCUt.', 'user', '2026-01-13 03:02:39'),
(8, 'dama', 'dama@gmail.com', '$2a$10$hXc63/Twmov97SA8qFUSOe7.gEOvwrSNbYotRyyGK9LTp.4.aiq/C', 'admin', '2026-01-13 03:29:00');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `nis` (`username`),
  ADD UNIQUE KEY `email` (`email`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
