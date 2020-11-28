from jinja2 import Environment, FileSystemLoader


def build_all():
    env = Environment(loader=FileSystemLoader("."), trim_blocks=True)
    tmpl = env.get_template("map_literal_test.gen.go.jinja")
    lengths = [
        3, 5, 9, 17, 33, 65, 129, 257, 513
    ]
    print(tmpl.render(lengths=lengths))


if __name__ == "__main__":
    build_all()