<?php
  $count = 100;
  $time_start = microtime(true);

  $servername = "localhost";
  $username = "root";
  $password = 'my$ql';
  $dbname = "golang";

  // Create connection
  $conn = new mysqli($servername, $username, $password, $dbname);
  // Check connection
  if ($conn->connect_error) {
      die("Connection failed: " . $conn->connect_error);
  }


          for ($i = 1; $i <= $count; $i++) {
            $sql = "INSERT INTO user VALUES($i, 'fname', 'lname', 'addresss')";
            $result = $conn->query($sql);


              // //  fmt.Println("send message " , i)
              // $i*1000;
               echo ("message ". $i . PHP_EOL);

          }
          $time_end = microtime(true);
          $time = $time_end - $time_start;
          echo ($time);
     ?>
