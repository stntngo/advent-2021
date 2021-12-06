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
                   (map #(reduce + %)))]

    (->> (map - floor (next floor))
         (filter #(> 0 %))
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
(defn parse-bit-array [bits]
  (Integer/parseInt (str/join bits) 2))

(defn parse-bits [line]
  (str/split line #""))

(defn transpose [matrix]
  (apply map vector matrix))

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

(defn power-consumption [numbers]
  (let [gamma (->> numbers
                   transpose
                   (map #(reduce count-bit [0 0] %))
                   (map gamma?))
        epsilon (map flip gamma)]
    (* (parse-bit-array gamma) (parse-bit-array epsilon))))

(defn bit-filter [numbers idx f]
  (if (= 1 (count numbers))
    (first numbers)
    (let [counter (->> numbers
                       transpose
                       (map #(reduce count-bit [0 0] %)))
          target (-> counter
                     (nth idx)
                     f)]
      (-> (filter (fn [line] (-> line
                                 (nth idx)
                                 (= target)))
                  numbers)
          (recur (inc idx) f)))))

(defn life-support [numbers]
  (let [carbon (-> numbers
                   (bit-filter 0 (fn [[zero one]] (if (>= one zero) "1" "0")))
                   parse-bit-array)
        oxygen (-> numbers
                   (bit-filter 0 (fn [[zero one]] (if (<= zero one) "0" "1")))
                   parse-bit-array)]
    (* carbon oxygen)))

(defrecord Node [count zero one])

(defn weighted-trie [[head & tail] node]
  (let [node (if node
               (update-in node [:count] inc)
               (Node. 1 nil nil))]
    (case head
      "0" (update-in node [:zero] #(weighted-trie tail %))
      "1" (update-in node [:one] #(weighted-trie tail %))
      nil node)))

(defn trie-filter [node f]
  (when (and node (or (:zero node) (:one node)))
    (let [[val next] (f (:zero node) (:one node))]
      (conj (trie-filter next f) val))))

(defn oxygen-trie-filter [zero one]
  (cond
    (and zero one) (if (>= (:count one) (:count zero))
                     [1 one]
                     [0 zero])
    zero [0 zero]
    one [1 one]
    :else [nil nil]))

(defn carbon-trie-filter [zero one]
  (cond
    (and zero one) (if (<= (:count zero) (:count one))
                     [0 zero]
                     [1 one])
    zero [0 zero]
    one [1 one]
    :else [nil nil]))

(defn life-support-trie [numbers]
  (let [root (->> numbers
                  (reduce #(weighted-trie %2 %1) (Node. 0 nil nil)))
        carbon (-> root
                   (trie-filter carbon-trie-filter)
                   parse-bit-array)
        oxygen (-> root
                   (trie-filter oxygen-trie-filter)
                   parse-bit-array)]
    (* carbon oxygen)))

(defn day-three []
  (let [numbers (->> (input 3)
                     (read-file parse-bits))
        power (power-consumption numbers)
        life (life-support numbers)
        life-trie (life-support-trie numbers)]

    (println "Day Three")
    (print "Part One: ")
    (println power)
    (print "Part Two: ")
    (println life)
    (print "Part Two (Trie): ")
    (println life-trie)))

; Day Four
(defn parse-board [board]
  (->> board
       (map (fn [line]
              (as-> line v
                (str/split v #" ")
                (remove str/blank? v)
                (map #(Integer/parseInt %) v))))))

(defn parse-bingo [input]
  (let [[raw-numbers & raw-boards] (->> input
                                        str/split-lines
                                        (remove str/blank?))
        random (as-> raw-numbers nums
                 (str/split nums #",")
                 (map #(Integer/parseInt %) nums))
        boards  (->> raw-boards
                     (partition 5)
                     (map parse-board))]

    [[random nil #{}] boards]))

(defn draw-number [[pick & numbers] drawn]
  [numbers pick (conj drawn pick)])

(defn winner? [drawn board]
  (or (->> board
           (some #(every? drawn %)))
      (->> board
           transpose
           (some #(every? drawn %)))))

(defn score-board [last-pick drawn board]
  (->> board
       flatten
       (remove drawn)
       (reduce +)
       (* last-pick)))

(defn first-winner [[numbers last-pick drawn] boards]
  (if-let [winner (->> boards
                       (filter #(winner? drawn %))
                       first)]
    [last-pick drawn winner]
    (recur (draw-number numbers drawn) boards)))

(defn last-winner [[numbers last-pick drawn] boards]
  (if (every? #(winner? drawn %) boards)
    [last-pick drawn (first boards)]
    (recur
     (draw-number numbers drawn)
     (remove #(winner? drawn %) boards))))

(defn day-four []
  (let [[random boards] (->> (input 4)
                             slurp
                             parse-bingo)]
    (println "Day Four")
    (print "Part One: ")
    (println (apply score-board (first-winner random boards)))
    (print "Part Two: ")
    (println (apply score-board (last-winner random boards)))))

; Day Five
(defrecord Point [x y])

(defn parse-point [s]
  (as-> s s
    (str/split s #",")
    (mapv #(Integer/parseInt %) s)
    (apply ->Point s)))

(defn parse-line [s]
  (as-> s s
    (str/split s #" -> ")
    (map parse-point s)))

(defn line-type [{x1 :x y1 :y}
                 {x2 :x y2 :y}]
  (cond
    (= x1 x2) :vertical
    (= y1 y2) :horizontal
    :else :diagonal))

(defn xrange [x y]
  (if (> x y)
    (reverse (range y (inc x)))
    (range x (inc y))))

(defn line-points [{x1 :x y1 :y}
                   {x2 :x y2 :y}]
  (case (line-type (Point. x1 y1) (Point. x2 y2))
    :vertical (map #(->Point x1 %) (xrange y1 y2))
    :horizontal (map #(->Point % y1) (xrange x1 x2))
    :diagonal (map #(->Point %1 %2) (xrange x1 x2) (xrange y1 y2))))

(defn count-hot-spots [lines]
  (->> lines
       (map #(apply line-points %))
       flatten
       frequencies
       seq
       (remove (fn [[_ n]] (= 1 n)))
       count))

(defn day-five []
  (let [lines (->> (input 5)
                   (read-file parse-line))
        total (count-hot-spots lines)
        no-diag (->> lines
                     (remove #(= :diagonal (apply line-type %)))
                     count-hot-spots)]
    (println "Day Five")
    (println "Part One:" no-diag)
    (println "Part Two:" total)))

; Day Six
(defn parse-lantern-fish [s]
  (->> (str/split s #",")
       (map #(Integer/parseInt %))
       frequencies
       seq
       (reduce
        (fn [acc [idx value]]
          (update acc idx #(+ value %)))
        (apply vector (repeat 9 0)))))

(defn simulate-population [population days]
  (if (= 0 days)
    population
    (let [base (apply vector (repeat 9 0))
          [head & tail] population
          shifted (concat tail '(0))]
      (recur
       (as-> base b
         (update b 6 (constantly head))
         (update b 8 (constantly head))
         (mapv + b shifted))
       (dec days)))))

(defn day-six []
  (let [population (-> (input 6)
                       slurp
                       parse-lantern-fish)]
    (println "Day Six")
    (println "Part One:" (reduce + (simulate-population population 80)))
    (println "Part Two:" (reduce + (simulate-population population 256)))))

(defn -main []
  (day-one)
  (println "")
  (day-two)
  (println "")
  (day-three)
  (println "")
  (day-four)
  (println "")
  (day-five)
  (println "")
  (day-six))

