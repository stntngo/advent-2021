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
