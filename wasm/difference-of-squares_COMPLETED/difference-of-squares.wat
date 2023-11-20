(module
  (func $squareOfSum (export "squareOfSum") (param $max i32) (result i32)
    local.get $max
    local.get $max
    i32.const 1
    i32.add
    i32.mul
    i32.const 2
    i32.div_u
    local.tee $max
    local.get $max
    i32.mul
  )

  (func $sumOfSquares (export "sumOfSquares") (param $max i32) (result i32)
    local.get $max
    local.get $max
    i32.const 1
    i32.add
    i32.mul
    local.get $max
    i32.const 2
    i32.mul
    i32.const 1
    i32.add
    i32.mul
    i32.const 6
    i32.div_u
  )

  (func (export "difference") (param $max i32) (result i32)
    local.get $max
    call $squareOfSum
    local.get $max
    call $sumOfSquares
    i32.sub
  )
)
