(ns bird-watcher)

(def last-week
   [0 2 5 3 7 8 4]
  )

(defn today [birds]
   (last birds)
  )

(defn inc-bird [birds]
  (let [last (last birds)
        new-counts (vec (pop birds))]
    (conj new-counts (inc last)))
  )

(defn day-without-birds? [vec]
    (or 
        (some #(= % 0) vec)
         false
     )
  )

(defn n-days-count [birds n]
    (reduce + (take n birds))
  )

(defn busy-days [birds]
    (count (filter #(>= % 5) birds))
  )

(defn odd-week? [birds]
(every? #(or (= % 1) (= % 0)) birds)
  )
