(ns interest-is-interesting)

(defn abs "(abs n) is the absolute value of n" [n]
  (if (neg? n) (- n) n))

(defn interest-rate
  "Calculates the interest rate based on the specified balance.
  If the balance is negative, the interest rate is -3.213.
  If the balance is positive and less than 1000, the interest rate is 0.5.
  If the balance is positive and greater or equal to 1000 but less than 5000, the interest rate is 1.621.
  If the balance is positive and greater or equal to 5000, the interest rate is 2.475."
  [balance]
      (cond
        (neg? balance) -3.213
        (< balance 1000.0M) 0.5
        (< balance 5000.0M) 1.621
        :else 2.475
      )
  )


(defn annual-balance-update
  "Calculates the annual balance update based on the current balance and the interest rate.
  The interest rate is calculated using the interest-rate function.
  The annual balance update is calculated as the current balance multiplied by (1 + the interest rate).
  :param balance: the current balance
  :return: the annual balance update"
  [balance]
    (+
    (* (bigdec (interest-rate balance)) 0.01M (abs balance))
    balance))


(defn amount-to-donate
  "Calculates the amount to donate to charities based on the balance and the tax-free percentage that the government allows.
  The calculation is as follows:
  - If the balance is negative, the result is the floor of the balance multiplied by the tax-free percentage.
  - If the balance is positive, the result is the floor of two times the balance multiplied by the tax-free percentage.
  
  :param balance: The balance to donate (in monetary units)
  :param tax-free-percentage: The tax-free percentage allowed by the government
  :return: The amount to donate (in monetary units)"
  [balance tax-free-percentage]
    (if (pos? balance) (int (* tax-free-percentage 0.01 balance 2)) 0)
  )