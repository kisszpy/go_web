-- MySQL dump 10.13  Distrib 8.0.25, for macos11 (x86_64)
--
-- Host: 127.0.0.1    Database: go_db
-- ------------------------------------------------------
-- Server version	8.0.25

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `post`
--

DROP TABLE IF EXISTS `post`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `post` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `author` varchar(20) NOT NULL DEFAULT '',
  `title` varchar(64) NOT NULL DEFAULT '',
  `content` text NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `status` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=102 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `post`
--

LOCK TABLES `post` WRITE;
/*!40000 ALTER TABLE `post` DISABLE KEYS */;
INSERT INTO `post` VALUES (1,'吴克雷斯','乌克兰干俄罗斯','内容很精彩','2022-04-06 12:54:20','2022-04-06 12:54:20',0),(2,'吴克雷斯\0','乌克兰干俄罗斯\0','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(3,'吴克雷斯','乌克兰干俄罗斯','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(4,'吴克雷斯','乌克兰干俄罗斯','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(5,'吴克雷斯','乌克兰干俄罗斯','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(6,'吴克雷斯','乌克兰干俄罗斯','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(7,'吴克雷斯','乌克兰干俄罗斯','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(8,'吴克雷斯','乌克兰干俄罗斯','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(9,'吴克雷斯','乌克兰干俄罗斯','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(10,'吴克雷斯','乌克兰干俄罗斯','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(11,'吴克雷斯	','乌克兰干俄罗斯	','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(12,'吴克雷斯\n','乌克兰干俄罗斯\n','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(13,'吴克雷斯','乌克兰干俄罗斯','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(14,'吴克雷斯','乌克兰干俄罗斯','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(15,'吴克雷斯\r','乌克兰干俄罗斯\r','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(16,'吴克雷斯','乌克兰干俄罗斯','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(17,'吴克雷斯','乌克兰干俄罗斯','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(18,'吴克雷斯','乌克兰干俄罗斯','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(19,'吴克雷斯','乌克兰干俄罗斯','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(20,'吴克雷斯','乌克兰干俄罗斯','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(21,'吴克雷斯','乌克兰干俄罗斯','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(22,'吴克雷斯','乌克兰干俄罗斯','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(23,'吴克雷斯','乌克兰干俄罗斯','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(24,'吴克雷斯','乌克兰干俄罗斯','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(25,'吴克雷斯','乌克兰干俄罗斯','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(26,'吴克雷斯','乌克兰干俄罗斯','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(27,'吴克雷斯','乌克兰干俄罗斯','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(28,'吴克雷斯\Z','乌克兰干俄罗斯\Z','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(29,'吴克雷斯','乌克兰干俄罗斯','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(30,'吴克雷斯','乌克兰干俄罗斯','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(31,'吴克雷斯','乌克兰干俄罗斯','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(32,'吴克雷斯','乌克兰干俄罗斯','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(33,'吴克雷斯','乌克兰干俄罗斯','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(34,'吴克雷斯 ','乌克兰干俄罗斯 ','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(35,'吴克雷斯!','乌克兰干俄罗斯!','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(36,'吴克雷斯\"','乌克兰干俄罗斯\"','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(37,'吴克雷斯#','乌克兰干俄罗斯#','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(38,'吴克雷斯$','乌克兰干俄罗斯$','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(39,'吴克雷斯%','乌克兰干俄罗斯%','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(40,'吴克雷斯&','乌克兰干俄罗斯&','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(41,'吴克雷斯\'','乌克兰干俄罗斯\'','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(42,'吴克雷斯(','乌克兰干俄罗斯(','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(43,'吴克雷斯)','乌克兰干俄罗斯)','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(44,'吴克雷斯*','乌克兰干俄罗斯*','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(45,'吴克雷斯+','乌克兰干俄罗斯+','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(46,'吴克雷斯,','乌克兰干俄罗斯,','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(47,'吴克雷斯-','乌克兰干俄罗斯-','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(48,'吴克雷斯.','乌克兰干俄罗斯.','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(49,'吴克雷斯/','乌克兰干俄罗斯/','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(50,'吴克雷斯0','乌克兰干俄罗斯0','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(51,'吴克雷斯1','乌克兰干俄罗斯1','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(52,'吴克雷斯2','乌克兰干俄罗斯2','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(53,'吴克雷斯3','乌克兰干俄罗斯3','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(54,'吴克雷斯4','乌克兰干俄罗斯4','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(55,'吴克雷斯5','乌克兰干俄罗斯5','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(56,'吴克雷斯6','乌克兰干俄罗斯6','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(57,'吴克雷斯7','乌克兰干俄罗斯7','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(58,'吴克雷斯8','乌克兰干俄罗斯8','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(59,'吴克雷斯9','乌克兰干俄罗斯9','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(60,'吴克雷斯:','乌克兰干俄罗斯:','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(61,'吴克雷斯;','乌克兰干俄罗斯;','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(62,'吴克雷斯<','乌克兰干俄罗斯<','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(63,'吴克雷斯=','乌克兰干俄罗斯=','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(64,'吴克雷斯>','乌克兰干俄罗斯>','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(65,'吴克雷斯?','乌克兰干俄罗斯?','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(66,'吴克雷斯@','乌克兰干俄罗斯@','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(67,'吴克雷斯A','乌克兰干俄罗斯A','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(68,'吴克雷斯B','乌克兰干俄罗斯B','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(69,'吴克雷斯C','乌克兰干俄罗斯C','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(70,'吴克雷斯D','乌克兰干俄罗斯D','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(71,'吴克雷斯E','乌克兰干俄罗斯E','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(72,'吴克雷斯F','乌克兰干俄罗斯F','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(73,'吴克雷斯G','乌克兰干俄罗斯G','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(74,'吴克雷斯H','乌克兰干俄罗斯H','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(75,'吴克雷斯I','乌克兰干俄罗斯I','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(76,'吴克雷斯J','乌克兰干俄罗斯J','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(77,'吴克雷斯K','乌克兰干俄罗斯K','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(78,'吴克雷斯L','乌克兰干俄罗斯L','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(79,'吴克雷斯M','乌克兰干俄罗斯M','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(80,'吴克雷斯N','乌克兰干俄罗斯N','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(81,'吴克雷斯O','乌克兰干俄罗斯O','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(82,'吴克雷斯P','乌克兰干俄罗斯P','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(83,'吴克雷斯Q','乌克兰干俄罗斯Q','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(84,'吴克雷斯R','乌克兰干俄罗斯R','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(85,'吴克雷斯S','乌克兰干俄罗斯S','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(86,'吴克雷斯T','乌克兰干俄罗斯T','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(87,'吴克雷斯U','乌克兰干俄罗斯U','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(88,'吴克雷斯V','乌克兰干俄罗斯V','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(89,'吴克雷斯W','乌克兰干俄罗斯W','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(90,'吴克雷斯X','乌克兰干俄罗斯X','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(91,'吴克雷斯Y','乌克兰干俄罗斯Y','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(92,'吴克雷斯Z','乌克兰干俄罗斯Z','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(93,'吴克雷斯[','乌克兰干俄罗斯[','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(94,'吴克雷斯\\','乌克兰干俄罗斯\\','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(95,'吴克雷斯]','乌克兰干俄罗斯]','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(96,'吴克雷斯^','乌克兰干俄罗斯^','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(97,'吴克雷斯_','乌克兰干俄罗斯_','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(98,'吴克雷斯`','乌克兰干俄罗斯`','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(99,'吴克雷斯a','乌克兰干俄罗斯a','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(100,'吴克雷斯b','乌克兰干俄罗斯b','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0),(101,'吴克雷斯c','乌克兰干俄罗斯c','内容很精彩-----','2022-04-06 12:55:44','2022-04-06 12:55:44',0);
/*!40000 ALTER TABLE `post` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `products`
--

DROP TABLE IF EXISTS `products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `products` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext,
  `sku` longtext,
  `price` bigint DEFAULT NULL,
  `create_time` datetime(3) DEFAULT NULL,
  `update_time` datetime(3) DEFAULT NULL,
  `status` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_products_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
/*!40000 ALTER TABLE `products` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t_price`
--

DROP TABLE IF EXISTS `t_price`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_price` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `province` varchar(20) NOT NULL DEFAULT '' COMMENT '省',
  `city` varchar(20) NOT NULL DEFAULT '' COMMENT '市',
  `sku` varchar(32) NOT NULL DEFAULT '' COMMENT '商品sku',
  `sku_name` varchar(128) NOT NULL DEFAULT '' COMMENT '商品名称',
  `market_price` int NOT NULL DEFAULT '0' COMMENT '市场价格',
  `channel` varchar(32) NOT NULL DEFAULT '' COMMENT '渠道',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `status` int NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_price`
--

LOCK TABLES `t_price` WRITE;
/*!40000 ALTER TABLE `t_price` DISABLE KEYS */;
INSERT INTO `t_price` VALUES (1,'','','1002','牛肉',0,'','2022-04-15 13:24:16','2022-04-15 13:24:16',0),(2,'','','1002','牛肉',0,'','2022-04-15 14:36:44','2022-04-15 14:36:44',0),(3,'','','1002','牛肉',0,'','2022-04-16 14:03:12','2022-04-16 14:03:12',0),(4,'','','1002','牛肉',0,'','2022-04-16 14:03:15','2022-04-16 14:03:15',0),(5,'','','1002','牛肉',0,'','2022-04-16 14:03:20','2022-04-16 14:03:20',0);
/*!40000 ALTER TABLE `t_price` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t_resource`
--

DROP TABLE IF EXISTS `t_resource`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_resource` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '资源名称',
  `pid` bigint unsigned NOT NULL DEFAULT '0' COMMENT '父节点',
  `data` json DEFAULT NULL COMMENT '数据',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `status` int unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='系统资源表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_resource`
--

LOCK TABLES `t_resource` WRITE;
/*!40000 ALTER TABLE `t_resource` DISABLE KEYS */;
INSERT INTO `t_resource` VALUES (1,'系统管理',0,'{\"meta\": {\"icon\": \"el-icon-setting\", \"title\": \"系统管理\"}, \"name\": \"System\", \"path\": \"/system\", \"redirect\": \"/system\", \"component\": \"Layout\", \"alwaysShow\": true}','2022-04-11 04:10:46','2022-04-11 04:10:46',0),(2,'角色管理',1,'{\"meta\": {\"icon\": \"email\", \"title\": \"角色管理\"}, \"name\": \"Role\", \"path\": \"role\", \"component\": \"system/role\"}','2022-04-11 04:01:38','2022-04-11 04:01:38',0),(5,'用户管理',1,'{\"meta\": {\"icon\": \"user\", \"title\": \"用户管理\"}, \"name\": \"User\", \"path\": \"user\", \"component\": \"system/user\"}','2022-04-11 04:01:32','2022-04-11 04:01:32',0),(9,'菜单管理',1,'{\"meta\": {\"icon\": \"el-icon-menu\", \"title\": \"菜单管理\"}, \"name\": \"Menu\", \"path\": \"menu\", \"component\": \"system/menu\"}','2022-04-12 00:52:21','2022-04-12 00:52:21',0),(34,'系统图标',1,'{\"meta\": {\"icon\": \"lock\", \"title\": \"系统图标\"}, \"name\": \"系统图标\", \"path\": \"icon\", \"redirect\": \"\", \"component\": \"system/icon\"}','2022-04-14 15:47:14','2022-04-14 15:47:14',0),(35,'订单管理',0,'{\"meta\": {\"icon\": \"lock\", \"title\": \"订单管理\"}, \"name\": \"订单管理\", \"path\": \"/order\", \"redirect\": \"/order\", \"component\": \"Layout\", \"alwaysShow\": true}','2022-04-17 22:41:09','2022-04-17 22:41:09',0),(38,'我的销售',35,'{\"meta\": {\"icon\": \"lock\", \"title\": \"我的销售\"}, \"name\": \"我的销售\", \"path\": \"sale\", \"redirect\": \"\", \"component\": \"order/sale\", \"alwaysShow\": false}','2022-04-17 14:36:18','2022-04-17 14:36:18',0),(39,'我的采购',35,'{\"meta\": {\"icon\": \"lock\", \"title\": \"我的采购\"}, \"name\": \"我的采购\", \"path\": \"purchase\", \"redirect\": \"\", \"component\": \"order/purchase\", \"alwaysShow\": false}','2022-04-17 14:45:51','2022-04-17 14:45:51',0),(40,'商品管理',0,'{\"meta\": {\"icon\": \"lock\", \"title\": \"商品管理\"}, \"name\": \"商品管理\", \"path\": \"/product\", \"redirect\": \"\", \"component\": \"Layout\", \"alwaysShow\": true}','2022-04-17 14:54:23','2022-04-17 14:54:23',0),(42,'类目管理',40,'{\"meta\": {\"icon\": \"lock\", \"title\": \"类目管理\"}, \"name\": \"类目管理\", \"path\": \"category\", \"redirect\": \"\", \"component\": \"product/category\", \"alwaysShow\": false}','2022-04-17 14:55:22','2022-04-17 14:55:22',0),(43,'品牌管理',40,'{\"meta\": {\"icon\": \"lock\", \"title\": \"品牌管理\"}, \"name\": \"品牌管理\", \"path\": \"brand\", \"redirect\": \"\", \"component\": \"product/brand\", \"alwaysShow\": false}','2022-04-17 14:56:25','2022-04-17 14:56:25',0);
/*!40000 ALTER TABLE `t_resource` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t_role`
--

DROP TABLE IF EXISTS `t_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_role` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `role_name` varchar(20) NOT NULL DEFAULT '' COMMENT '角色名称',
  `role_desc` varchar(20) NOT NULL DEFAULT '' COMMENT '角色描述',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `status` int NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_role`
--

LOCK TABLES `t_role` WRITE;
/*!40000 ALTER TABLE `t_role` DISABLE KEYS */;
INSERT INTO `t_role` VALUES (1,'超级管理员','超级管理员','2022-04-13 16:36:09','2022-04-13 16:36:09',0),(6,'客服小组','客服小组000','2022-04-14 19:00:01','2022-04-14 19:00:01',0);
/*!40000 ALTER TABLE `t_role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t_role_permission`
--

DROP TABLE IF EXISTS `t_role_permission`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_role_permission` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `role_id` bigint NOT NULL DEFAULT '0' COMMENT '角色id',
  `resource_id` bigint NOT NULL DEFAULT '0' COMMENT '资源id',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `status` int NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色权限表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_role_permission`
--

LOCK TABLES `t_role_permission` WRITE;
/*!40000 ALTER TABLE `t_role_permission` DISABLE KEYS */;
INSERT INTO `t_role_permission` VALUES (1,1,1,'2022-04-17 14:50:17','2022-04-17 14:50:17',0),(2,1,2,'2022-04-17 14:50:32','2022-04-17 14:50:32',0),(3,1,5,'2022-04-17 14:50:38','2022-04-17 14:50:38',0),(4,1,9,'2022-04-17 14:50:45','2022-04-17 14:50:45',0),(5,1,34,'2022-04-17 14:50:51','2022-04-17 14:50:51',0),(13,1,35,'2022-04-17 14:29:40','2022-04-17 14:29:40',0);
/*!40000 ALTER TABLE `t_role_permission` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t_test`
--

DROP TABLE IF EXISTS `t_test`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_test` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `balance` int NOT NULL DEFAULT '0',
  `frozen` int NOT NULL DEFAULT '0',
  `available` int NOT NULL DEFAULT '0',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `status` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_test`
--

LOCK TABLES `t_test` WRITE;
/*!40000 ALTER TABLE `t_test` DISABLE KEYS */;
INSERT INTO `t_test` VALUES (1,20000,380,19620,'2022-04-16 23:23:30','2022-04-16 23:23:30',0);
/*!40000 ALTER TABLE `t_test` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t_user`
--

DROP TABLE IF EXISTS `t_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(20) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(20) NOT NULL DEFAULT '' COMMENT '密码',
  `email` varchar(50) NOT NULL DEFAULT '' COMMENT '邮箱',
  `nickname` varchar(50) NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar` varchar(512) DEFAULT '' COMMENT '头像',
  `introduction` varchar(200) DEFAULT '' COMMENT '简介',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `status` int NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='系统用户表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_user`
--

LOCK TABLES `t_user` WRITE;
/*!40000 ALTER TABLE `t_user` DISABLE KEYS */;
INSERT INTO `t_user` VALUES (1,'admin','000000','guaga@163.com','孤星一族','','','2022-04-17 15:03:13','2022-04-17 15:03:13',0),(25,'baba','000000000','baba@163.com','baba','','','2022-04-17 13:02:49','2022-04-17 13:02:49',0),(26,'keting','00000000','keting@000.com','keting','','','2022-04-17 14:57:53','2022-04-17 14:57:53',0);
/*!40000 ALTER TABLE `t_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t_user_role`
--

DROP TABLE IF EXISTS `t_user_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_user_role` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL DEFAULT '0' COMMENT '用户id',
  `role_id` bigint NOT NULL DEFAULT '0' COMMENT '角色id',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `status` int NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t_user_role`
--

LOCK TABLES `t_user_role` WRITE;
/*!40000 ALTER TABLE `t_user_role` DISABLE KEYS */;
INSERT INTO `t_user_role` VALUES (1,1,1,'2022-04-13 16:49:33','2022-04-13 16:49:33',0);
/*!40000 ALTER TABLE `t_user_role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_info`
--

DROP TABLE IF EXISTS `user_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_info` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '名称',
  `password` varchar(20) NOT NULL DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_info`
--

LOCK TABLES `user_info` WRITE;
/*!40000 ALTER TABLE `user_info` DISABLE KEYS */;
INSERT INTO `user_info` VALUES (1,'jack','000000'),(2,'lisi','aaa'),(3,'lisi','aaa'),(4,'lisi','999999'),(5,'lisi','999999'),(6,'java','pwd0000'),(7,'java','pwd0000'),(8,'java','pwd0000'),(9,'java','pwd0000'),(10,'java','pwd0000'),(11,'java','pwd0000'),(12,'java','pwd0000'),(13,'java','pwd0000'),(14,'java','pwd0000'),(15,'java','pwd0000'),(16,'java','pwd0000'),(17,'java','pwd0000'),(18,'java','pwd0000'),(19,'java','pwd0000'),(20,'java','pwd0000'),(21,'java','pwd0000');
/*!40000 ALTER TABLE `user_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_info_v2`
--

DROP TABLE IF EXISTS `user_info_v2`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_info_v2` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(20) NOT NULL DEFAULT '',
  `password` varchar(32) NOT NULL DEFAULT '',
  `email` varchar(20) NOT NULL DEFAULT '',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `status` int unsigned DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_info_v2`
--

LOCK TABLES `user_info_v2` WRITE;
/*!40000 ALTER TABLE `user_info_v2` DISABLE KEYS */;
INSERT INTO `user_info_v2` VALUES (1,'email@999999','java','java@java.com','2022-04-07 01:35:33','2022-04-07 01:35:33',0),(4,'jack','','kisszpy@999.com','2022-04-08 06:33:06','2022-04-08 06:33:06',0),(5,'jack','e10adc3949ba59abbe56e057f20f883e','kisszpy@999.com','2022-04-08 06:38:17','2022-04-08 06:38:17',0);
/*!40000 ALTER TABLE `user_info_v2` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `username` varchar(20) NOT NULL,
  `usernick` varchar(50) NOT NULL,
  `phone` varchar(11) NOT NULL,
  `password` varchar(255) NOT NULL,
  `loginip` varchar(20) NOT NULL,
  `email` varchar(255) NOT NULL,
  `isadmin` tinyint(1) DEFAULT NULL,
  `nickname` varchar(50) NOT NULL,
  `login_ip` varchar(20) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `phone` (`phone`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-04-18 14:02:40
