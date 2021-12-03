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

; Day Three
(defn parse-bits [line]
  (str/split line, #""))

(defn transpose [matrix]
  (apply mapv vector matrix))

(defn count-bit [[zero one] bit]
  (case bit
    "0" [(inc zero) one]
    "1" [zero (inc one)]))

(defn gamma? [[zero one]]
  (if (> one zero) "1" "0"))

(defn flip [bit]
  (case bit
    "0" "1"
    "1" "0"))

(defn parse-bit-array [bits]
  (Integer/parseInt (str/join bits) 2))

(defn power-consumption [numbers]
  (let [gamma (->> numbers
                   transpose
                   (mapv (partial reduce count-bit [0 0]))
                   (mapv gamma?))
        epsilon (mapv flip gamma)]
    (* (parse-bit-array gamma) (parse-bit-array epsilon))))

(defn day-three []
  (let [numbers (->> (input 3)
                     (read-file parse-bits))
        power (power-consumption numbers)]
    (println "Day Three")
    (print "Part One: ")
    (println power)))

(defn main []
  ; (day-one)
  ; (day-two)
  (day-three))



