components {
  id: "player"
  component: "/main/player.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"Spaceship\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/game.atlas\"\n"
  "}\n"
  ""
  position {
    x: 2.0
    y: -3.0
  }
  scale {
    x: 0.101724
    y: 0.113524
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"player\"\n"
  "mask: \"star\"\n"
  "mask: \"hazard\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 3.0\n"
  "      y: -4.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "    id: \"Player Basket\"\n"
  "  }\n"
  "  data: 81.52117\n"
  "  data: 95.9189\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "beam_factory"
  type: "factory"
  data: "prototype: \"/main/beam.go\"\n"
  ""
}
embedded_components {
  id: "sfx_beam"
  type: "sound"
  data: "sound: \"/main/beam.wav\"\n"
  ""
}
