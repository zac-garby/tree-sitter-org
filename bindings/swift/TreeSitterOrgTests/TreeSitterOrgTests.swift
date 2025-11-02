import XCTest
import SwiftTreeSitter
import TreeSitterOrg

final class TreeSitterOrgTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_org())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading Org mode grammar")
    }
}
