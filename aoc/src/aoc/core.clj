(ns aoc.core
  (:gen-class)
  (:require [clojure.java.io :as io])
  (:require [clojure.string :as str]))

(defn read-file [p f]
  (with-open [rdr (io/reader f)]
    (->> rdr
         line-seq
         (mapv p))))

(defn input [day]
  (str "/Users/niels/code/stntngo/advent/" day "/input"))

; Day One
(defn sweep-floor [readings window]
  (let [floor (->> readings
                   (partition window 1)
                   (map (partial reduce +)))]

    (->> (map - floor (next floor))
         (filter (partial > 0))
         count)))

(defn day-one []
  (let [sweep (->> (input 1)
                   (read-file #(Integer/parseInt %))
                   (partial sweep-floor))]
    (println "Day One")
    (print "Part One: ")
    (println (sweep 1))
    (print "Part Two: ")
    (println (sweep 3))))

; Day Two
(defn parse-command [line]
  (let [[dir dist] (str/split line #" ")]
    (vector (case dir
              "up" :up
              "down" :down
              "forward" :forward)
            (Integer/parseInt dist))))

(defn aim [[x y aim] [direction distance]]
  (case direction
    :forward [(+ x distance) (+ y (* aim distance)) aim]
    :up [x y (- aim distance)]
    :down [x y (+ aim distance)]))

(defn day-two []
  (let [[x y aim] (->> (input 2)
                       (read-file parse-command)
                       (reduce aim [0 0 0]))]
    (println "Day Two")
    (print "Part One: ")
    (println (* x aim))
    (print "Part Two: ")
    (println (* x y))))

(defn main []
  (day-one)
  (day-two))

