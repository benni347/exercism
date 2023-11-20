(module
  (func (export "squareRoot") (param $radicand i32) (result f32)
    local.get $radicand
    f32.convert_i32_s
    f32.sqrt)
)
