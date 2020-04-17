import setuptools

with open("README.md", "r") as fh:
    long_description = fh.read()

setuptools.setup(
    name="mackerel-project_validator",
    version="0.0.1",
    author="Mario Lubenka",
    author_email="me@saitho.me",
    description="Validate your Mackerel project definition files",
    long_description=long_description,
    long_description_content_type="text/markdown",
    url="https://github.com/mackerelserver/project-validator",
    packages=['mackerel_project_validator'],
    package_data={'mackerel_project_validator': ['bin/*', 'schema/*']},
    classifiers=[
        "Environment :: Console",
        "Intended Audience :: Developers",
        "Intended Audience :: System Administrators",
        "License :: OSI Approved :: GNU General Public License v2 or later (GPLv2+)",
        "Operating System :: OS Independent",
        "Topic :: Software Development :: Testing",
        "Topic :: Software Development :: Quality Assurance",
    ]
)
