<?php
  $count = 100000000;
  $time_start = microtime(true);
  $current = 0;
  $current2 = 0;


          for ($i = 1; $i <= $count; $i++) {
            $current += $i*2/10*(22/12*(12/22));
            $current2 += $i*2/10*(22/12*(12/22));
          }

          $time_end = microtime(true);
          $time = $time_end - $time_start;
          echo ($time." ".$current);
     ?>
