<?php
  $count = 200;
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
            $sql = "select u.* from users u left join address a on u.id = a.u_id   WHERE u.id >  $i";
            $result = $conn->query($sql);


                // output data of each row
                while($row = $result->fetch_assoc()) {
                    // $temp =   $row["id"]. " " . $row["fname"]. " " . $row["lname"]. PHP_EOL;
                }

              // //  fmt.Println("send message " , i)
              // $i*1000;
              //  echo ("message ". $i . PHP_EOL);

          }
          $time_end = microtime(true);
          $time = $time_end - $time_start;
          echo ($time);
     ?>
