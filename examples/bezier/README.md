# Bezier rotation animation

```Go
	ctx := gog.NewContext(250, 250)
	bezierPath := shapes.CubicBezier(100, 95, 50, 300, 190, 88, 140, 200, 50)
	bezierPath.SetPos(ctx.Center).Scale(vec.Vec2{1.3, 1.3})
	for i := 0; i < 150; i++ {
		ctx.Clear(color.Gray{30})
		bezierPath.Rotate((math.Pi * 2) / 150)
		ctx.DebugDraw(bezierPath)
		ctx.AppendAnimationFrame()
	}
	ctx.SaveAPNG("bezier_anim.png", 3)
```

![bezier anim](https://github.com/user-attachments/assets/135882e0-5a6d-438c-b0d0-80ebd12713c2)
