# Curve animation example

```Go
  c := gog.NewContext(250, 250)
  curve := shapes.CubicBezier(100, 95, 50, 300, 190, 88, 140, 200, 50)
  curve.SetPos(c.Center).Scale(vec.Vec2{1.3, 1.3})
  for i := 0; i < 150; i++ {
      c.Clear(color.Gray{30})
      curve.Rotate((math.Pi * 2) / 150)
      c.DebugDraw(curve)
      c.AppendAnimationFrame()
  }
  c.SaveAPNG("curve_anim.png", 3)
```

![curve_anim](https://github.com/user-attachments/assets/135882e0-5a6d-438c-b0d0-80ebd12713c2)
