""" Management of the schema """

from flask_script import Manager
from flask_migrate import MigrateCommand
from kingpin.application import application

manager = Manager(application)
manager.add_command('db', MigrateCommand)


def main():
    """ Allows an entrypoint for management of the schema """
    manager.run()


if __name__ == '__main__':
    main()
