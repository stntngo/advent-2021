(ns aoc.core-test
  (:require [clojure.test :as test]
            [clojure.string :as str]
            [aoc.core :as aoc]))

(defn read-str [s p]
  (doall (map p (str/split-lines s))))

(def day-one-test-case "199
200
208
210
200
207
240
269
260
263")

(def day-two-test-case "forward 5
down 5
forward 8
up 3
down 8
forward 2")

(def day-three-test-case "00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010")

(def day-four-test-case "7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7")

(test/deftest day-one-parse-test
  (test/testing "Day One Parse Test"
    (test/is (= 10 (count (read-str day-one-test-case #(Integer/parseInt %)))))))

(test/deftest day-one-depth-test
  (test/testing "Day One Window Test"
    (test/is (= 7 (aoc/sweep-floor
                   (read-str day-one-test-case #(Integer/parseInt %))
                   1)))))

(test/deftest day-one-window-test
  (test/testing "Day One Window Test"
    (test/is (= 5 (aoc/sweep-floor
                   (read-str day-one-test-case #(Integer/parseInt %))
                   3)))))

(test/deftest day-two-parse-test
  (test/testing "Day Two Parse Test"
    (test/is (= 6 (count
                   (read-str day-two-test-case aoc/parse-command))))))

(test/deftest day-two-aim-test
  (test/testing "Day Two Aim Test"
    (let [[x y aim] (reduce aoc/aim [0 0 0] (read-str day-two-test-case aoc/parse-command))]
      (test/is (= 150 (* x aim)))
      (test/is (= 900 (* x y))))))

(test/deftest day-three-power-test
  (test/testing "Day Three Power Test"
    (test/is (= 198 (-> day-three-test-case
                        (read-str aoc/parse-bits)
                        aoc/power-consumption)))))

(test/deftest day-three-life-support-test
  (test/testing "Day Three Life Support Test"
    (test/is (= 230 (-> day-three-test-case
                        (read-str aoc/parse-bits)
                        aoc/life-support)))))

(test/deftest day-four-parse-test
  (test/testing "Day Four Parse Test"
    (let [[random boards] (-> day-four-test-case
                               aoc/parse-bingo)
          [_ drawn _] (apply aoc/draw-number random)]
      (test/is (= 3 (count boards)))
      (test/is (= 7 drawn)))))

(test/deftest day-four-winners
  (test/testing "Day Four Winner Test"
    (let [[random boards] (-> day-four-test-case
                               aoc/parse-bingo)]
      (test/is (= 4512 (apply aoc/score-board (aoc/first-winner random boards))))
      (test/is (= 1924 (apply aoc/score-board (aoc/last-winner random boards)))))))
