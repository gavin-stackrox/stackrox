import scopeSelectors from '../helpers/scopeSelectors';
import search from '../selectors/search';

export const selectors = {
    globalSearchButton: 'button:contains("Search")',
    pageSearch: scopeSelectors('[data-testid="page-header"]', {
        input: search.input,
        options: search.input,
    }),
    globalSearch: scopeSelectors('.search-modal', {
        input: search.input,
        options: search.input,
    }),
    empty: scopeSelectors('.pf-v5-c-empty-state', {
        head: 'h1',
        body: '.pf-v5-c-empty-state__body',
    }),
    tab: 'li.pf-v5-c-tabs__item',
    count: '.pf-v5-c-badge',
    globalSearchResults: scopeSelectors('[data-testid="global-search-results"]', {
        header: 'h1',
    }),

    // Include ancestor selector like section#All to match only the table for the active tab.
    viewOnChip: 'td[data-label="View On:"] button',
    filterOnChip: 'td[data-label="Filter On:"] button',
};

export default selectors;
