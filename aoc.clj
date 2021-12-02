(ns aoc
  (:gen-class)
  (:require [clojure.java.io :as io])
  (:require [clojure.string :as str]))

(defn read-file [f p]
  (with-open [rdr (io/reader f)]
    (doall (map p (lazy-seq (line-seq rdr))))))

; Day One
(defn sweep-floor [readings window]
  (let [floor (map (partial reduce +) (partition window 1 readings))]
    (count
     (filter (partial apply <) (map vector floor (drop 1 floor))))))

(let [sweep (partial sweep-floor (read-file "1/input" #(Integer/parseInt %)))]
  (println "Day One")
  (print "Part One: ")
  (println (sweep 1))
  (print "Part Two: ")
  (println (sweep 3)))

; Day Two
(defn parse-command [line]
  (let [[dir dist] (str/split line #" ")]
    (list (case dir
            "up" :up
            "down" :down
            "forward" :forward)
          (Integer/parseInt dist))))

(defn aim [[x y aim] [direction distance]]
  (case direction
    :forward [(+ x distance) (+ y (* aim distance)) aim]
    :up [x y (- aim distance)]
    :down [x y (+ aim distance)]))

(let [[x y aim] (reduce aim [0 0 0] (read-file "2/input" parse-command))]
  (println "Day Two")
  (print "Part One: ")
  (println (* x aim))
  (print "Part Two: ")
  (println (* x y)))
